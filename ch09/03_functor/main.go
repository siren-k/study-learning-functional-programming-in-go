package main

import (
	"encoding/json"
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch09/03_functor/functor"
)

var (
	ClockUnit = func(i int) int {
		return i
	}

	AmPmMapper = func(i int) int {
		return (i + 12) % 24
	}
)

func AmHoursFn() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
}

func main() {
	fmt.Println("----------------------------------------")
	fmt.Println("- Clock Functor")
	fmt.Println("----------------------------------------")
	fmt.Println("initial state    : ", NewClockFunctor(AmHoursFn()))
	fmt.Println("unit application : ", NewClockFunctor(AmHoursFn()).Map(ClockUnit))
	fmt.Println("1st application  : ", NewClockFunctor(AmHoursFn()).Map(AmPmMapper))
	fmt.Println("chain application: ", NewClockFunctor(AmHoursFn()).Map(AmPmMapper).Map(AmPmMapper))
	fmt.Println()

	fmt.Println("----------------------------------------")
	fmt.Println("- Car Functor")
	fmt.Println("----------------------------------------")
	cars := []Car{
		{"Honda", "Accord"},
		{"Lexus", "IS250"},
	}
	str := `{"make": "Toyota", "model": "Highlander"}`
	highlander := Car{}
	json.Unmarshal([]byte(str), &highlander)
	cars = append(cars, highlander)

	fmt.Println("initial state   : ", CarWrap(cars))
	fmt.Println("unit application: ", CarWrap(cars).Map(CarUnit))
	fmt.Println("one upgrade     : ", CarWrap(cars).Map(CarUpgrade))
	fmt.Println("chain upgrades  : ", CarWrap(cars).Map(CarUpgrade).Map(CarUpgrade))
	fmt.Println("one downgrade   : ", CarWrap([]Car{{"Honda", "Accord"}, {"Lexus", "IS250 LX"}, {"Toyota", "Highlander LX Limited"}}).Map(CarDowngrade))
	fmt.Println("up and downgrade: ", CarWrap(cars).Map(CarUpgrade).Map(CarDowngrade))
}
