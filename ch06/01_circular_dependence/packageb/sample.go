package packageb

import (
	"fmt"
	a "github.com/siren-k/study-learning-functional-programming-in-go/ch06/01_circular_dependence/packagea"
)

func Btask() {
	fmt.Println("B")
	a.Atask()
}
