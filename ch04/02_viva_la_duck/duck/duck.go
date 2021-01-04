package duck

import (
	"errors"
	"fmt"
	"log"
)

// The Pond structure contains information about the pond,
// that is, how many insects a duck can eat and how many
// strokes it takes to traverse it.
type Pond struct {
	BugSupply      int
	StrokeRequired int
}

// One of the first things to do is to define the system
// behavior in a simple interface. In addition, it should
// be considered that the behavior pattern of the system
// consists of a set of basic interfaces. Since things
// can be defined through their actions, it is reasonable
// to classify things according to actions.
//
// The main benefit of an interface is that it categorizes
// the functions of an application so that it can model
// behavior in the real world.
type StrokeBehavior interface {
	PaddleFoot(strokeSupply *int)
}

type EatBehavior interface {
	EatBug(strokeSupply *int)
}

// Declaring a small, single-purpose interface allows you
// to use it freely when defining a new interface that
// provides rich functionality.
type SurvivalBehaviors interface {
	StrokeBehavior
	EatBehavior
}

type Duck struct{}

// Will you pass a value or a reference?
// The first rule of thumb is to pass a reference (pointer type)
// if you want to share the state. If not, pass the value. In the
// Stroke function, it is necessary to update the strokeSupply
// type of the duck, so it is passed in the form of an integer
// pointer. Pointer parameters are passed only when absolutely
// necessary. It's time to write code defensively. Because someone
// can try to run code in parallel. Passing parameters by value
// is safe for parallel execution. When passing a reference, it
// is necessary to define a critical section using sync.Mutex.
func (Duck) Stroke(s StrokeBehavior, strokeSupply *int, p Pond) (err error) {
	for i := 0; i < p.StrokeRequired; i++ {
		if *strokeSupply < p.StrokeRequired-1 {
			err = errors.New("Our duck died!")
		}
		s.PaddleFoot(strokeSupply)
	}
	return err
}

type Foot struct{}

func (Foot) PaddleFoot(strokeSupply *int) {
	fmt.Println("- Foot, paddle!")
	*strokeSupply--
}

func (Duck) Eat(e EatBehavior, strokeSupply *int, p Pond) {
	for i := 0; i < p.BugSupply; i++ {
		e.EatBug(strokeSupply)
	}
}

type Bill struct{}

func (Bill) EatBug(strokeSupply *int) {
	*strokeSupply++
	fmt.Println("- Bill, eat a bug!")
}

// SwimAndEat takes StrokeAndEatBehaviors, a set of interfaces,
// and uses them to express Stroke and Eat polymorphically.
func (d Duck) SwimAndEat(se SurvivalBehaviors, strokeSupply *int, ponds []Pond) {
	for i := range ponds {
		pond := &ponds[i]
		// d.Stroke function takes SurvivalBehaviors type as if
		// it was passed StrokeBehavior.
		err := d.Stroke(se, strokeSupply, *pond)
		if err != nil {
			log.Fatal(err)
		}
		// d.Eat function takes SurvivalBehaviors type as if
		// it was passed EatBehavior.
		d.Eat(se, strokeSupply, *pond)
	}
}
