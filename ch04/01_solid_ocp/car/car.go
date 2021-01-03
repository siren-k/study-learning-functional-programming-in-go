package car

import "fmt"

type Car struct {
	Make  string
	Model string
}

func (car Car) Tires() int {
	return 4
}

func (c Car) PrintInfo() {
	fmt.Printf("%v has %d tires\n", c, c.Tires())
}

type CarWithSpare struct {
	Car
}

func (car CarWithSpare) Tires() int {
	return 5
}
