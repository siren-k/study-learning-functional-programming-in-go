package functor

type IntFunctor interface {
	Map(f func(int) int) IntFunctor
}

type intBox struct {
	ints []int
}

func (box intBox) Map(f func(int) int) IntFunctor {
	for i, el := range box.ints {
		box.ints[i] = f(el)
	}
	return box
}
