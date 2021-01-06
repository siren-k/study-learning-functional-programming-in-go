package packagea

import (
	"fmt"
	b "github.com/siren-k/study-learning-functional-programming-in-go/ch06/01_circular_dependence/packageb"
)

func Atask() {
	fmt.Println("A")
	b.Btask()
}
