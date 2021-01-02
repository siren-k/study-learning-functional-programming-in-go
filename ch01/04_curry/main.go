package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch01/04_curry/internal"
)

// Function literals are functions that are treated as first-class objects of
// the language. First-class objects are elements that become variable types,
// such as int or string. Go functions can be declared as types. In addition,
// it can be assigned to a field of a variable or structure, and can be a
// parameter or return value of another function. Function literals can access
// the scope of the environment declared as a closure. When a function literal
// is assigned to a variable at execution time like
// val := func(x int) int {return x + 2 }(5), the anonymous function at this
// time can be called a function expression. Function literals are used with
// currying in lambda expressions.
func main() {
	fn1 := CurryAddTwo
	fn2 := CurryAddThree

	fmt.Printf("CurryAddTwo(%d) returned %d\n", 1, fn1(1))
	fmt.Printf("CurryAddTwo(%d) returned %d\n", 2, fn1(2))

	fmt.Printf("CurryAddThree(%d) returned %d\n", 1, fn2(1))
	fmt.Printf("CurryAddThree(%d) returned %d\n", 2, fn2(2))
}
