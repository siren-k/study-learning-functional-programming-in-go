package internal

type fibonacci func(int) int

// Before the main() function is executed, the memoized() function is executed
// and the returned anonymous function is assigned to the memoizedFibonacci variable.
var memoizedFibonacci = momoized(Fibonacci)

func momoized(fib fibonacci) fibonacci {
	// Map type is assigned to the variable cache for caching Fibonacci calculation values.
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := fib(key)
		cache[key] = temp
		return temp
	}
}

// Memoization is an optimization method that improves performance by storing
// the result of a function call that consumes a lot of resources and then
// returning the saved result value when the same input is given. Memoization
// works without problems due to the two properties of pure functions.
// * A pure function always returns the same result if the input is the same.
// * Does not cause side effects on the execution environment.
func MemoizedFibonacci(n int) int {
	return memoizedFibonacci(n)
}
