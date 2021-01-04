package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch04/02_viva_la_duck/duck"
)

const DASHES = "----------------------------------------"

type Capabilities struct {
	StrokeBehavior
	EatBehavior
	strokes int
}

func displayDuckStats(c Capabilities, ponds []Pond) {
	fmt.Printf("%s\n", DASHES)
	fmt.Printf("Ponds Processed:")
	for _, pond := range ponds {
		fmt.Printf("\n\t%+v", pond)
	}
	fmt.Printf("\nStrokes remaining: %+v\n", c.strokes)
	fmt.Printf("%s\n", DASHES)
}

// The lessons of this story are as follows.
// * Model your application like the real world in a meaningful way.
// * Create a set of actions with a single task interface type.
// * Maintains simple interface types consistently and merges them into extended behavior sets.
// * Make sure that each function takes only the required behavior type.
// * Let's not be a duck.
func main() {
	var duck Duck
	Capabilities := Capabilities{
		StrokeBehavior: Foot{},
		EatBehavior:    Bill{},
		strokes:        5,
	}

	ponds := []Pond{
		{BugSupply: 1, StrokeRequired: 3},
		{BugSupply: 1, StrokeRequired: 2},
	}
	duck.SwimAndEat(&Capabilities, &Capabilities.strokes, ponds)
	displayDuckStats(Capabilities, ponds)

	ponds = []Pond{
		{BugSupply: 2, StrokeRequired: 3},
	}
	duck.SwimAndEat(&Capabilities, &Capabilities.strokes, ponds)
	displayDuckStats(Capabilities, ponds)
}
