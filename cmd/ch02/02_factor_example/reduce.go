package main

import (
	"fmt"
	. "github.com/wesovilabs/koazee"
)

func main() {
	numbers := []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

	fmt.Printf("input: %v\n", numbers)
	stream := StreamOf(numbers)

	fmt.Print("stream.Reduce(sum): ")
	fmt.Println(stream.Reduce(func(acc, val int) int {
		return acc + val
	}).Int())

	// ❯ go run reduce.go
	// input: [1 5 4 3 2 7 1 8 2 3]
	// stream.Reduce(sum): 36
}
