package functor

import "fmt"

type ClockFunctor interface {
	Map(f func(int) int) ClockFunctor
}

type hourContainer struct {
	hours []int
}

func (box hourContainer) Map(f func(int) int) ClockFunctor {
	for i, el := range box.hours {
		box.hours[i] = f(el)
	}
	return box
}

func (box hourContainer) String() string {
	return fmt.Sprintf("%+v", box.hours)
}

func NewClockFunctor(hours []int) ClockFunctor {
	return hourContainer{hours: hours}
}
