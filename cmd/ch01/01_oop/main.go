package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/cmd/ch01/01_oop/internal"
)

func main() {
	MyCars.Add(Car{"IS250"})
	MyCars.Add(Car{"Blazer"})
	MyCars.Add(Car{"Highlander"})

	car, err := MyCars.Find("Highlander")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Found : %v\n", car)
	}

	car, err = MyCars.Find("ighlander")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Found : %v\n", car)
	}
}
