package internal

func CurryAddTwo(n int) (ret int) {
	defer func() {
		ret = n + 2
	}()
	return n
}

func CurryAddThree(n int) int {
	// func(n int) int {return n + 3 }(5) is an anonymous function,
	// a function literal, a closure, and a lambda expression.
	// Function literals are
	// * It is written like a function declaration, but there is no name following the func keyword.
	// * It is a formula.
	// * Can access variables within the grammatical scope.
	return func(n int) int {
		return n + 3
	}(n)
}
