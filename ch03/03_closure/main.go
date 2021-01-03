package main

import "fmt"

var sum = 5

// The closure consists of the addTwo function. In addTwo,
// sum and anonymous functions are declared within the same
// lexicographic scope. Since addTwo has scope to cover sum
// and anonymous functions, and sum is declared before
// anonymous functions, anonymous functions can always access
// and modify the sum variable. As soon as addTwo is assigned
// to twoMore, the anonymous function of the addTwo function
// accesses the sum variable and persists for the duration
// of application execution.
func addTwo() func() int {
	sum := 0
	return func() int {
		sum += 2
		return sum
	}
}

// Even if the stack frame is deepened by several layers of
// addTwoDynamics, the Go runtime searches the environment
// where addTwoDynamic is defined for sum search. If it is
// not found at the defined location, it continues up the
// stack and searches until a sum can be found. Therefore,
// the dynamic range increases complexity, and the value of
// sum may change in unexpected ways or patterns that are
// difficult to debug.
func addTwoDynamic() func() int {
	return func() int {
		sum += 2
		return sum
	}
}

// Being able to pass a lexical context is a powerful feature,
// and ensures that it does not involve side effects that may
// occur in the dynamic scope.
func main() {
	twoMore := addTwo()
	fmt.Println(twoMore())
	fmt.Println(twoMore())

	twoMoreDynamic := addTwoDynamic()
	fmt.Println(twoMoreDynamic())
	fmt.Println(twoMoreDynamic())
}
