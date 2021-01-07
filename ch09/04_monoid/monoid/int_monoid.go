package monoid

type IntMonoid interface {
	Append(i ...int) IntMonoid
	Zero() []int
	Reduce() int
}

type intContainer struct {
	Ints []int
}

func (box intContainer) Append(ints ...int) IntMonoid {
	box.Ints = append(box.Ints, ints...)
	return box
}

func (box intContainer) Zero() []int {
	return nil
}

func (box intContainer) Reduce() int {
	total := 0
	for _, item := range box.Ints {
		total += item
	}
	return total
}

func WrapInt(ints []int) IntMonoid {
	return intContainer{Ints: ints}
}
