package main

import (
	"fmt"
	"math/rand"
)

type Stock interface {
	register(observer StockObserver)
	unregister(observer StockObserver)
	notifyAll()
	getStockPrice() float64
}

type StockObserver interface {
	update(float64)
	getUserName() string
}

type StockMarket struct {
	observers  []StockObserver
	stockPrice float64
}

type StockTrader struct {
	username string
}

func (market *StockMarket) register(observer StockObserver) {
	market.observers = append(market.observers, observer)
}

func (market *StockMarket) unregister(observer StockObserver) {
	index := -1
	for i, obs := range market.observers {
		if obs.getUserName() == observer.getUserName() {
			index = i
			break
		}
	}

	if index != -1 {
		market.observers = append(market.observers[:index], market.observers[index+1:]...)
	}
}

func (market *StockMarket) notifyAll() {
	for _, observer := range market.observers {
		observer.update(market.stockPrice)
	}
}

func (market *StockMarket) getStockPrice() float64 {
	return market.stockPrice
}

func (trader *StockTrader) update(price float64) {
	fmt.Printf("%s, Price updated: $%.2f\n", trader.username, price)
}

func (trader *StockTrader) getUserName() string {
	return trader.username
}

func main() {
	stockMarket := &StockMarket{}

	trader1 := &StockTrader{username: "Trader1"}
	trader2 := &StockTrader{username: "Trader2"}
	trader3 := &StockTrader{username: "Trader3"}

	stockMarket.register(trader1)
	stockMarket.register(trader2)
	stockMarket.register(trader3)

	initialPrice := 50.0
	stockMarket.stockPrice = initialPrice
	fmt.Printf("\nPrice updated: $%.2f\n", initialPrice)
	stockMarket.notifyAll()

	for i := 0; i < 3; i++ {
		newPrice := initialPrice + float64(i+rand.Int())
		stockMarket.stockPrice = newPrice
		fmt.Printf("\nPrice updated: $%.2f\n", newPrice)
		stockMarket.notifyAll()
	}
}
