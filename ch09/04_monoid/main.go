package main

import (
	"fmt"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch09/04_monoid/monoid"
)

func main() {
	const name = "Alice"
	stringMonoid := monoid.WrapName(name)
	fmt.Println("----------------------------------------")
	fmt.Println("- NameMonoid")
	fmt.Println("----------------------------------------")
	fmt.Println("Initial state    : [", stringMonoid, "]")
	fmt.Println("Zero             : [", stringMonoid.Zero(), "]")
	fmt.Println("1st application  : [", stringMonoid.Append(name), "]")
	fmt.Println("Chain application: [", stringMonoid.Append(name).Append(name), "]")

	ints := []int{1, 2, 3}
	intMonoid := monoid.WrapInt(ints)
	fmt.Println("----------------------------------------")
	fmt.Println("- IntMonoid")
	fmt.Println("----------------------------------------")
	fmt.Println("Initial state    : [", intMonoid, "]")
	fmt.Println("Zero             : [", intMonoid.Zero(), "]")
	fmt.Println("1st application  : [", intMonoid.Append(ints...), "]")
	fmt.Println("Chain application: [", intMonoid.Append(ints...).Append(ints...), "]")
	fmt.Println("Reduce chain     : [", intMonoid.Append(ints...).Append(ints...).Reduce(), "]")

	lineItems := []monoid.LineItem{
		{Quantity: 1, Price: 12978, ListPrice: 22330},
		{Quantity: 2, Price: 530, ListPrice: 786},
		{Quantity: 5, Price: 270, ListPrice: 507},
	}
	lineItemMonoid := monoid.WrapLineItem(lineItems)
	fmt.Println("----------------------------------------")
	fmt.Println("- LineItemMonoid")
	fmt.Println("----------------------------------------")
	fmt.Println("Initial state    : [", lineItemMonoid, "]")
	fmt.Println("Zero             : [", lineItemMonoid.Zero(), "]")
	fmt.Println("1st application  : [", lineItemMonoid.Append(lineItems...), "]")
	fmt.Println("Chain application: [", lineItemMonoid.Append(lineItems...).Append(lineItems...), "]")
	fmt.Println("Reduce chain     : [", lineItemMonoid.Append(lineItems...).Append(lineItems...).Reduce(), "]")
}
