package gogomatcher

import (
	// "errors"
	"github.com/google/uuid"
	// "math"
)

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

// func MatchOrder(exchange *Exchange, order *Order) (bool, error) {
// 	orderbook := *exchange.orderbooks[order.ticker]
// 	levelsToSearch := orderbook.asks

// 	// for searching through bids
// 	bidLimitCheck := func(levelPrice float32, limit float32) bool {
// 		return levelPrice >= limit
// 	}

// 	// for searching through asks
// 	askLimitCheck := func(levelPrice float32, limit float32) bool {
// 		return levelPrice <= limit
// 	}

// 	limitCheck := askLimitCheck

// 	switch order.side {
// 	case Bid:
// 		levelsToSearch = orderbook.asks
// 		limitCheck = askLimitCheck
// 	case Ask:
// 		levelsToSearch = orderbook.bids
// 		limitCheck = bidLimitCheck
// 	default:
// 		return false, errors.New("not a valid order side")
// 	}

// 	priceLimit := math.Inf(1)

// 	if order.orderType == Limit {
// 		switch order.side {
// 		case Bid:
// 			priceLimit = math.MaxFloat32
// 		case Ask:
// 			priceLimit = 0 // TODO: can this be negative?
// 		default:
// 			return false, errors.New("not a valid order side")
// 		}
// 	}

// 	for idx, level := range levelsToSearch {
// 		if !(limitCheck(level.price, priceLimit)) {
// 			break
// 		}
// 	}

// 	return true, nil
// }
