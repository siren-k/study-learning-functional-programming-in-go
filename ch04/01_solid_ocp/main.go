package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch04/01_solid_ocp/car"
)

// Go supports:
// * If not defined, the method of the built-in object is implicitly called.
// * Manually calling another object's method
// * Overriding methods of built-in objects
func main() {
	accord := Car{"Honda", "Accord"}
	accord.PrintInfo()

	highlander := CarWithSpare{Car{"Toyota", "Highlander"}}
	highlander.PrintInfo()

	fmt.Printf("%v has %d tires\n", highlander.Car, highlander.Tires())
}
