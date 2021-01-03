package main

import (
	"fmt"
	"strings"
)

type StrFunc func(string) string

func Compose(f StrFunc, g StrFunc) StrFunc {
	return func(s string) string {
		return g(f(s))
	}
}

// We can combine the functions f() and g() to get from A to B
// and from B to C. Note that the order is important. First of all,
// go from A to B using the function f(), and go from B to C
// using the function g(). This can be expressed as (f.g)(x).
// This is read as'f-Compose-g's input x'. This formula is
// equivalent to g(f(x)). That is, (f.g)(x) == g(f(x)).
func main() {
	var recognize = func(name string) string {
		return fmt.Sprintf("Hey %s", name)
	}
	var emphasize = func(statement string) string {
		return fmt.Sprintf(strings.ToUpper(statement) + "!")
	}

	var greetFoG = Compose(recognize, emphasize)
	fmt.Println(greetFoG("Gopher"))

	var greetGoF = Compose(emphasize, recognize)
	fmt.Println(greetGoF("Gopher"))
}
