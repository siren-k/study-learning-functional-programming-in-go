package main

import (
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch02/01_collection_interation/internal"
)

// There is a problem that the desired and implementation methods are mixed.
func main() {
	collection := NewSlice([]string{"CRV", "IS250", "Blazer"})
	value, ok := collection.Next()
	// In functional programming, we use a method of declaring what to do
	// rather than instructing the implementation in detail step by step.
	for ok {
		println(value)
		value, ok = collection.Next()
	}
}
