package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch01/03_sum/internal"
)

func main() {
	fmt.Println(SumLoop([]int{1, 2, 3}))
	fmt.Println(SumRecursive([]int{1, 2, 3}))
}
