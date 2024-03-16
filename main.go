package main

import (
	"fmt"
	"github.com/google/uuid"
	"sort"
)

type OrderBook struct {
	BuyOrders  []Order
	SellOrders []Order
}

type Order struct {
	OrderID   uuid.UUID
	BuyOrSell string
	Price     float64
	Quantity  int32
}

func (ob *OrderBook) AddOrder(order Order) {
	if order.Price < 999_999 && order.Quantity < 999_999_999 {
		if order.BuyOrSell == "B" {
			ob.BuyOrders = append(ob.BuyOrders, order)

			sort.Slice(ob.BuyOrders, func(i, j int) bool {
				return ob.BuyOrders[i].Price < ob.BuyOrders[j].Price
			})

		} else if order.BuyOrSell == "S" {
			ob.SellOrders = append(ob.SellOrders, order)
			sort.Slice(ob.SellOrders, func(i, j int) bool {
				return ob.SellOrders[i].Price > ob.SellOrders[j].Price
			})
		} else {
			fmt.Println(order, " its uncorrected order")

		}

	} else {
		fmt.Println("uncorrected order")

	}
}

func (ob *OrderBook) RemoveOrder(order Order) {
	if order.BuyOrSell == "B" {
		for i, o := range ob.BuyOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.BuyOrders = append(ob.BuyOrders[:i], ob.BuyOrders[i+1:]...)

				break
			}
		}
	} else if order.BuyOrSell == "S" {
		for i, o := range ob.SellOrders {
			if o.Price == order.Price && o.Quantity == order.Quantity {
				ob.SellOrders = append(ob.SellOrders[:i], ob.SellOrders[i+1:]...)
				break
			}
		}
	}

}

func (ob *OrderBook) MatchOrders() {
	for len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0 {
		buyOrder := ob.BuyOrders[0]
		sellOrder := ob.SellOrders[0]
		if buyOrder.Price == sellOrder.Price {

			ob.RemoveOrder(buyOrder)
			ob.RemoveOrder(sellOrder)
			if buyOrder.Quantity > sellOrder.Quantity {
				remainOrder := Order{OrderID: uuid.New(),
					BuyOrSell: "B",
					Price:     buyOrder.Price,
					Quantity:  buyOrder.Quantity - sellOrder.Quantity}
				ob.AddOrder(remainOrder)

			}
			fmt.Printf("Matched order , delete this buyOrder(%v) & this sellOrder(%v)\n", buyOrder, sellOrder)
		} else {
			break
		}

	}

}

func main() {

	orderBook := OrderBook{}
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "B", Price: 1000, Quantity: 100})
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "S", Price: 1000, Quantity: 50})
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "S", Price: 1000, Quantity: 50})
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "S", Price: 1000, Quantity: 50})
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "S", Price: 1000, Quantity: 50})
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "S", Price: 1000, Quantity: 50})
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "falseOeder", Price: 1000, Quantity: 50})
	orderBook.AddOrder(Order{OrderID: uuid.New(), BuyOrSell: "S", Price: 1000, Quantity: 50})

	orderBook.MatchOrders()

	fmt.Printf("%v\n%v\n", orderBook.BuyOrders, orderBook.SellOrders)

}
