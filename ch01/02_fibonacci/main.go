package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch01/02_fibonacci/internal"
)

func main() {
	fmt.Println("simple fibonacci: ", Fibonacci(8))
	fmt.Println("memoized fibonacci: ", MemoizedFibonacci(8))
	fmt.Println("channeled fibonacci: ", ChanneledFibonacci(8))
}
