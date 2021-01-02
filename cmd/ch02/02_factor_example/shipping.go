package main

import (
	"fmt"
	"github.com/siren-k/study-learning-functional-programming-in-go/cmd/ch02/02_factor_example/cart"
	. "github.com/wesovilabs/koazee"
	"log"
	"strings"
	"time"
)

const discount = 0.8

type ItemTotal struct {
	*cart.Item
	total float32
}

var shippingStream = Stream().
	Filter(func(item *cart.Item) bool {
		return item.ExpirationDate.After(time.Now())
	}).
	Sort(func(itemLeft, itemRight *cart.Item) int {
		return strings.Compare(itemLeft.Name, itemRight.Name)
	}).
	Map(func(item *cart.Item) *ItemTotal {
		total := float32(item.Units) * item.PricePerUnit
		if item.Units >= 3 {
			total *= discount
		}
		return &ItemTotal{item, total}
	}).
	ForEach(func(item *ItemTotal) {
		fmt.Printf(" - %s, %d units, %.2f€\n", item.Name, item.Units, item.total)
	})

func process(items []*cart.Item) {
	myStream := shippingStream.With(items).Do()
	output :=
		myStream.
			Reduce(func(acc float32, item *ItemTotal) float32 {
				return acc + item.total
			})
	if output.Err() != nil {
		log.Fatal(output.Err().Error())
	}
	fmt.Printf(" Total price %.2f€\n", output.Float32())
}

func main() {
	fmt.Println(":--------------------------:")
	process(cart.Items)
	fmt.Println(":--------------------------:")
	process(cart.Items2)
	fmt.Println(":--------------------------:")

	// ❯ go run shipping.go
	// :--------------------------:
	//  - Avocado, 7 units, 5.88€
	//  - Eggs Free range, 2 units, 4.50€
	//  - Milk, 4 units, 3.94€
	//  - Onions, 5 units, 1.08€
	//  Total price 15.40€
	// :--------------------------:
	//  - Carrots, 9 units, 1.22€
	//  - Cucumber, 3 units, 0.60€
	//  - Kale, 1 units, 1.23€
	//  - Potatoes, 7 units, 2.52€
	//  - Tofu, 1 units, 3.05€
	//  Total price 8.62€
}
