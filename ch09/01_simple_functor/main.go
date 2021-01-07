package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	ints := []int{1, 2, 3}

	imperativeInts := []int{}
	for _, v := range ints {
		imperativeInts = append(imperativeInts, v+1)
	}
	fmt.Println("imperative loop:", imperativeInts)

	add1 := func(i int) int { return i + 1 }
	fpInts := koazee.StreamOf(ints).Map(add1).Do().Out().Val()
	fmt.Println("fp map:", fpInts)
}
