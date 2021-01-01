package main

import (
	"fmt"
	. "github.com/siren-k/study-learning-functional-programming-in-go/cmd/ch01/01_oop/internal"
)

// 본질적으로 명령행 코드와 동일한 연산을 수행한다. 프로그램의 상태를 객체의 속성에 배정하고,
// 메소드를 호출해 내부 상태를 수정하고, 원하는 결과에 도달할 때까지 실행 상태를 변경한다.
func main() {
	MyCars.Add(Car{"IS250"})
	MyCars.Add(Car{"Blazer"})
	MyCars.Add(Car{"Highlander"})

	car, err := MyCars.Find("Highlander")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Found : %v\n", car)
	}

	car, err = MyCars.Find("ighlander")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Found : %v\n", car)
	}
}
