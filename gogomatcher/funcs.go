package gogomatcher

import (
	"errors"
	"math"
	"slices"

	"github.com/google/uuid"
)

/*	TODO LIST:
	- Insertions for bid levels work but for ask levels it is janked - I think I may need to go back to the drawing board
	and change how the levels are structured - but first we need to try to make a somewhat clean solution for the current
	implementation to see if it is worth switching or not.
*/

func NewAccount(username string) *Account {
	account := new(Account)
	account.uid = uuid.New()
	account.username = username

	return account
}

func AddNewAccount(exchange *Exchange, username string) uuid.UUID {
	account := NewAccount(username)
	exchange.accounts[account.uid] = account

	return account.uid
}

func GetAccount(exchange *Exchange, uuid uuid.UUID) *Account {
	return exchange.accounts[uuid]
}

// type Order struct {
// 	uid       uuid.UUID
// 	ticker    string
// 	price     float32
// 	quantity  float32
// 	side      Side
// 	orderType OrderType
// }

func NewOrder(ticker string, price float32, quantity float32, side Side, orderType OrderType) *Order {
    order := new(Order)
	order.uid = uuid.New()
	order.ticker = ticker
	order.price = price
	order.quantity = quantity
	order.side = side
	order.orderType = orderType

	return order
}

func GetOrderBook(exchange Exchange, ticker string) *OrderBook {
	return exchange.orderbooks[ticker]
}

func NewOrderBook(ticker string) *OrderBook {
	orderbook := new(OrderBook)
	orderbook.ticker = ticker
	orderbook.bids = make([]Level, 0)
	orderbook.asks = make([]Level, 0)
	return orderbook
}

func NewExchange() *Exchange {
	exchange := new(Exchange)
	exchange.orderbooks = make(map[string]*OrderBook)
	exchange.accounts = make(map[uuid.UUID]*Account)
	return exchange
}

func AddOrderBookToExchange(exchange *Exchange, orderbook *OrderBook) {
	_, ok := exchange.orderbooks[orderbook.ticker]
	if !ok {
		exchange.orderbooks[orderbook.ticker] = orderbook
	}
}

func NewLevel(price float32) *Level {
	level := new(Level)
	level.price = price
	level.quantity = 0.0
	level.orders = make([]*Order, 0)

	return level
}

func MatchOrder(exchange *Exchange, order *Order) (bool, error) {
	orderbook := exchange.orderbooks[order.ticker]
	levelsToSearch := &orderbook.asks

	// for searching through bids
	bidLimitCheck := func(levelPrice float32, limit float32) bool {
		return levelPrice >= limit
	}

	// for searching through asks
	askLimitCheck := func(levelPrice float32, limit float32) bool {
		return levelPrice <= limit
	}

	limitCheck := askLimitCheck

	switch order.side {
	case Bid:
		levelsToSearch = &orderbook.asks
		limitCheck = askLimitCheck
	case Ask:
		levelsToSearch = &orderbook.bids
		limitCheck = bidLimitCheck
	default:
		return false, errors.New("not a valid order side")
	}

	priceLimit := float32(0.0) // TODO: better way to do this?

	if order.orderType == Limit {
		switch order.side {
		case Bid:
			priceLimit = math.MaxFloat32
		case Ask:
			priceLimit = 0 // TODO: can this be negative?
		default:
			return false, errors.New("not a valid order type")
		}
	}

	// search for level, then orders to match with
	for _, level := range *levelsToSearch {
		if !(limitCheck(level.price, priceLimit)) {
			break
		}
	}

	switch order.side {
	case Bid:
		levelsToSearch = &orderbook.bids
		limitCheck = bidLimitCheck
	case Ask:
		levelsToSearch = &orderbook.asks
		limitCheck = askLimitCheck
	default:
		return false, errors.New("not a valid order side")
	}

	toInsertAt := 0
	found := false;
	currIdx := 0;

	// search for level to insert order into if not (fully) matched
	for idx, level := range *levelsToSearch {
		currIdx++;
		if level.price == order.price {
			(*levelsToSearch)[idx].orders = append((*levelsToSearch)[idx].orders, order)
			return true, nil
		} else if !limitCheck(level.price, order.price) {
            if idx > 1 {
                toInsertAt = idx - 1
				found = true;
            } 
            break
		}
	}

	// we have to append the level to the slice of the slice
	if len(*levelsToSearch) > 0 && !found && currIdx == len(*levelsToSearch){
		newLevel := NewLevel(order.price)
		newLevel.orders = append(newLevel.orders, order)
		*levelsToSearch = append(*levelsToSearch, *newLevel)
		return true, nil
	}

	// appropriate price level does not exist, so we have to create it and then insert the order into it
	newLevel := NewLevel(order.price)
	newLevel.orders = append(newLevel.orders, order)
	*levelsToSearch = slices.Insert(*levelsToSearch, toInsertAt, *newLevel)

	return true, nil
}
