package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch09/02_gen/car"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch09/02_gen/employee"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch09/02_gen/num"
)

func main() {
	fmt.Println("----------------------------------------")
	fmt.Println("- CarSlice Example")
	fmt.Println("----------------------------------------")
	var cars = CarSlice{
		Car{"Honda", "Accord", 3000},
		Car{"Lexus", "IS250", 40000},
		Car{"Toyota", "Highlander", 3500},
		Car{"Honda", "Accord ES", 3500},
	}
	fmt.Printf("cars: %+v\n", cars)

	honda := func(c Car) bool {
		return c.Make == "Honda"
	}
	fmt.Printf("filter cars by 'Honda': %+v\n", cars.Where(honda))

	price := func(c Car) Dollars {
		return c.Price
	}
	fmt.Printf("Hondas prices: %+v\n", cars.Where(honda).SelectDollars(price))

	fmt.Printf("Hondas sum(prices): %+v\n", cars.Where(honda).SumDollars(price))

	// The gen tool saves you the trouble of typing a lot of repetitive boilerplate text.
	fmt.Println("----------------------------------------")
	fmt.Println("- Sum Example")
	fmt.Println("----------------------------------------")
	fmt.Println("int8sum:", Int8Slice{1, 2, 3}.SumInt8(Int8fn))
	fmt.Println("int32sum:", Int32Slice{1, 2, 3}.SumInt32(Int32fn))
	fmt.Println("float64sum:", Float64Slice{1, 2, 3}.SumFloat64(Float64fn))
	fmt.Println("complex128sum:", Complex128Slice{1, 2, 3}.SumComplex128(Complex128fn))

	fmt.Println("----------------------------------------")
	fmt.Println("- Aggregate Example")
	fmt.Println("----------------------------------------")
	employees := EmployeeSlice{
		{Name: "Alice", Department: "Accounting"},
		{Name: "Bob", Department: "Back Office"},
		{Name: "Carly", Department: "Containers"},
	}
	join := func(previous string, e Employee) string {
		if previous != "" {
			previous += ", "
		}
		return previous + e.Name
	}
	fmt.Printf("aggregate: %+v\n", employees.AggregateString(join))
}
