package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

			} else if sellOrder.Quantity > buyOrder.Quantity {
				remainOrder := Order{OrderID: uuid.New(),
					BuyOrSell: "S",
					Price:     sellOrder.Price,
					Quantity:  sellOrder.Quantity - buyOrder.Quantity}
				ob.AddOrder(remainOrder)

			}
			fmt.Printf("Matched order , delete this buyOrder(%v) & this sellOrder(%v)\n", buyOrder, sellOrder)

		} else {
			break
		}

	}

}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
	return string(s)
}

func InputFromUser() {
	reader := bufio.NewReader(os.Stdin)
	var orderBook OrderBook

	for {
		var order Order
		fmt.Println("please enter your order in this format (B/S,Price,Quantity)")
		orderInput, _ := reader.ReadString('\n')
		orderParts := strings.Split(strings.TrimSpace(orderInput), ",")
		if len(orderParts) != 3 {
			fmt.Println("your order is not valid  please try again")
			continue
		}
		orderId, _ := uuid.NewRandom()
		order.OrderID = orderId

		order.BuyOrSell = orderParts[0]

		price, err := strconv.ParseFloat(orderParts[1], 64)
		if err != nil {
			log.Fatal(err)

		}
		order.Price = price
		quantity, err := strconv.Atoi(orderParts[2])
		if err != nil {
			log.Fatal(err)

		}
		order.Quantity = int32(quantity)

		orderBook.AddOrder(order)
		orderBook.MatchOrders()
		fmt.Println(orderBook)
		{
			fmt.Println("Do you want to add another order?  (y/n)")
			choiceInput, _ := reader.ReadString('\n')
			choice := strings.TrimSpace(choiceInput)

			if choice == "y" {
				continue
			} else if choice == "n" {
				fmt.Println("do you want to print all and pretty the orderBook?(y/n")
				PrintOrderChoiceInput, _ := reader.ReadString('\n')
				PrintOrderChoice := strings.TrimSpace(PrintOrderChoiceInput)
				if PrintOrderChoice == "y" {
					prettyPrint(orderBook)
					break
				}
				break
			}
		}

	}
}

func main() {
	InputFromUser()
}
