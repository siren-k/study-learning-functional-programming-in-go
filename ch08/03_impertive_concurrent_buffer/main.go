package main

import (
	"fmt"
	gc "github.com/go-goodies/go_currency"
)

type Order struct {
	OrderNumber     int
	IsAuthenticated bool
	IsDecrypted     bool
	Credentials     string
	CCardNumber     string
	CCardExpDate    string
	LineItems       []LineItem
}

type LineItem struct {
	Description string
	Count       int
	PriceUSD    gc.USD
}

func GetOrders() []*Order {
	order1 := &Order{
		OrderNumber:     10001,
		IsAuthenticated: true,
		IsDecrypted:     true,
		Credentials:     "alice,secret",
		CCardNumber:     "7b/HWvtIB9a16AYk+Yv6WWwer3GFbxpjoR+GO9iHIYY=",
		CCardExpDate:    "0922",
		LineItems: []LineItem{
			LineItem{"Apples", 1, gc.USD{4, 50}},
			LineItem{"Oranges", 4, gc.USD{12, 00}},
		},
	}

	order2 := &Order{
		OrderNumber:     10002,
		IsAuthenticated: true,
		IsDecrypted:     true,
		Credentials:     "bob,secret",
		CCardNumber:     "EOc3kF/OmxY+dRCaYRrey8h24QoGzVU0/T2QKVCHb1Q=",
		CCardExpDate:    "0123",
		LineItems: []LineItem{
			LineItem{"Milk", 2, gc.USD{8, 00}},
			LineItem{"Sugar", 1, gc.USD{2, 25}},
			LineItem{"Salt", 3, gc.USD{3, 75}},
		},
	}

	orders := []*Order{order1, order2}
	return orders
}

func Pipeline(o Order) Order {
	o = Authenticate(o)
	o = Decrypt(o)
	o = Charge(o)
	return o
}

func Authenticate(o Order) Order {
	fmt.Printf("Order %d is Authenticated\n", o.OrderNumber)
	return o
}

func Decrypt(o Order) Order {
	fmt.Printf("Order %d is Decrypted\n", o.OrderNumber)
	return o
}

func Charge(o Order) Order {
	fmt.Printf("Order %d is Charged\n", o.OrderNumber)
	return o
}

func main() {
	orders := GetOrders()
	numberOfOrders := len(orders)
	input := make(chan Order, numberOfOrders)
	output := make(chan Order, numberOfOrders)
	for i := 0; i < numberOfOrders; i++ {
		go func() {
			for order := range input {
				output <- Pipeline(order)
			}
		}()
	}

	for _, order := range orders {
		input <- *order
	}
	close(input)

	for i := 0; i < numberOfOrders; i++ {
		fmt.Println("The result is:", <-output)
	}
}
