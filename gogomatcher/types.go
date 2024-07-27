package gogomatcher

import "github.com/google/uuid"

type Side int8
type OrderType int8

const (
	Bid Side = iota
	Ask
)

const (
	Market OrderType = iota
	Limit
)

type Order struct {
	uid       uuid.UUID
	ticker    string
	price     float32
	quantity  float32
	side      Side
	orderType OrderType
}

type Level struct {
	price    float32
	quantity float32
	orders   []*Order
}

type OrderBook struct {
	ticker string
	bids   []Level
	asks   []Level
}

type StockTrack struct {
	ticker   string
	quantity float32
}

type Portfolio struct {
}

type Account struct {
	uid       uuid.UUID
	username string
	balance  float32
	stocks   map[string]StockTrack
}

type Exchange struct {
	orderbooks map[string]*OrderBook
	accounts   map[uuid.UUID]*Account
}
