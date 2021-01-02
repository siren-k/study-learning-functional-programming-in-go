package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch01/01_oop/internal"
)

// Essentially, it performs the same operations as command line code.
// It assigns the state of the program to the properties of the object,
// calls methods to modify the internal state, and changes the execution
// state until the desired result is reached.
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
