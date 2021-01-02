package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestSumRecursive(t *testing.T) {
	g := Goblin(t)
	g.Describe("Sum Recursive", func() {
		// Passing Test
		v := SumRecursive([]int{1, 2, 3})
		g.It("Should be 6", func() {
			g.Assert(v).Equal(6)
		})

		// Failing Test
		v = SumRecursive([]int{1, 2, 3})
		g.It("Should be 5", func() {
			g.Assert(v).Equal(5)
		})

		// Pending Test
		g.It("Should be 4")

		// Excluded Test
		g.Xit("Should be 3", func() {
			g.Assert(v).Equal(3)
		})
	})
}
