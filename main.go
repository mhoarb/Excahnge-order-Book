package main

import (
	"fmt"
	"sort"
)

type OrderBook struct {
	BuyOrders  []Order
	SellOrders []Order
}

type Order struct {
	IsBuy    bool //if they wish Buy or sell
	Price    float64
	Quantity int32 //the number of shares
}

func (ob *OrderBook) AddOrder(order Order) {
	if order.Price < 999_999 && order.Quantity < 999_999_999 {
		if order.IsBuy {
			ob.BuyOrders = append(ob.BuyOrders, order)

			sort.Slice(ob.BuyOrders, func(i, j int) bool {
				return ob.BuyOrders[i].Price < ob.BuyOrders[j].Price
			})

		} else {
			ob.SellOrders = append(ob.SellOrders, order)
			sort.Slice(ob.SellOrders, func(i, j int) bool {
				return ob.SellOrders[i].Price > ob.SellOrders[j].Price
			})
		}
		fmt.Println(ob)
		// fmt.Printf("OrderBook(BUY) :%v\n ,OrderBook(SELL) : %v " , ob.BuyOrders , ob.SellOrders)
	} else {
		fmt.Println("uncorrected order")

	}
}

func (ob *OrderBook) RemoveOrder(order Order) {
	if order.IsBuy {
		for i, o := range ob.BuyOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.BuyOrders = append(ob.BuyOrders[:i], ob.BuyOrders[i+1:]...)

			}

		}
	} else {
		for i, o := range ob.SellOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.SellOrders = append(ob.SellOrders[:i], ob.BuyOrders[i+1:]...)
				break
			}
		}
	}

}

func (ob *OrderBook) MatchOrders() {
	for len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0 {
		buyOrder := ob.BuyOrders[0]
		sellOrder := ob.SellOrders[0]
		if buyOrder.Price >= sellOrder.Price {
			ob.RemoveOrder(buyOrder)
			ob.RemoveOrder(sellOrder)
			fmt.Printf("Matched order , delete this buyOrder(%v) & this sellOrder(%v)", buyOrder, sellOrder)
		} else {
			break
		}

	}

}

func main() {
	orderBook := OrderBook{}
	orderBook.AddOrder(Order{Price: 10, Quantity: 100, IsBuy: true})
	orderBook.AddOrder(Order{Price: 9, Quantity: 50, IsBuy: true})
	orderBook.AddOrder(Order{Price: 11, Quantity: 200, IsBuy: true})
	orderBook.AddOrder(Order{Price: 10, Quantity: 50, IsBuy: false})

}
