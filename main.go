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

	// new order
	order0 := gogomatcher.NewOrder("TSLA", 69.42, 20, gogomatcher.Bid, gogomatcher.Market)
	order1 := gogomatcher.NewOrder("TSLA", 68.42, 20, gogomatcher.Bid, gogomatcher.Market)
	order2 := gogomatcher.NewOrder("TSLA", 70.42, 20, gogomatcher.Bid, gogomatcher.Market)
	order6 := gogomatcher.NewOrder("TSLA", 71.42, 20, gogomatcher.Bid, gogomatcher.Market)
	order3 := gogomatcher.NewOrder("TSLA", 80.42, 20, gogomatcher.Ask, gogomatcher.Market)
	order4 := gogomatcher.NewOrder("TSLA", 79.42, 20, gogomatcher.Ask, gogomatcher.Market)
	order5 := gogomatcher.NewOrder("TSLA", 81.42, 20, gogomatcher.Ask, gogomatcher.Market)

	_, err := gogomatcher.MatchOrder(exchange, order0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))

	_, err = gogomatcher.MatchOrder(exchange, order1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))

	_, err = gogomatcher.MatchOrder(exchange, order2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))

	_, err = gogomatcher.MatchOrder(exchange, order3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))

	_, err = gogomatcher.MatchOrder(exchange, order4)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))

	_, err = gogomatcher.MatchOrder(exchange, order5)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))

	_, err = gogomatcher.MatchOrder(exchange, order6)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("TSLA orderbook: %+v\n", gogomatcher.GetOrderBook(*exchange, "TSLA"))
}
