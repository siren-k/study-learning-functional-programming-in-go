package main

import "fmt"

// numberIs is a simple function that takes an integer
// and returns a boolean value
type numberIs func(int) bool

func lessThanTwo(i int) bool {
	return i < 2
}

func lessThan(x, y int) bool {
	return x < y
}

func (f numberIs) apply(s []int) (ret []bool) {
	for _, i := range s {
		ret = append(ret, f(i))
	}
	return ret
}

func main() {
	fmt.Println("NonCurried - lessThan(1, 2):", lessThan(1, 2))
	fmt.Println("Curried - lessThanTwo(1):", lessThanTwo(1))

	isLessThanOne := numberIs(func(i int) bool { return i < 1 }).apply
	isLessThanTwo := numberIs(lessThanTwo).apply
	s := []int{0, 1, 2}
	fmt.Println("Curried, given:", s, "...")
	fmt.Println("isLessThanOne:", isLessThanOne(s))
	fmt.Println("isLessThanTwo:", isLessThanTwo(s))
}
