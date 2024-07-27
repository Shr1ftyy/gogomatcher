package main

import (
	"fmt"
	"gogomatcher/gogomatcher"
)

func main() {
	fmt.Println("gogomatcher goooo!!")
	orderbook := gogomatcher.NewOrderBook("TSLA")
	exchange := gogomatcher.NewExchange()

    acctUuid := gogomatcher.AddNewAccount(exchange, "mikeburry")
    
	gogomatcher.AddOrderBookToExchange(exchange, orderbook)
	fmt.Printf("exchange: %+v\n", exchange)
	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))
	fmt.Printf("account: %+v\n", gogomatcher.GetAccount(exchange, acctUuid))
}
