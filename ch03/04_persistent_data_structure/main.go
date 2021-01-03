package main

import (
	"fmt"
	"github.com/mndrix/ps"
)

func main() {
	list1 := ps.NewList()
	fmt.Println(list1)

	list2 := list1.Cons("test1")
	list2.ForEach(func(v interface{}) {
		fmt.Printf("list2: %v, size: %d\n", v, list2.Size())
	})

	list3 := list2.Cons("test2")
	list3.ForEach(func(v interface{}) {
		fmt.Printf("list3: %v, size: %d\n", v, list3.Size())
	})
}
