package internal

// Instruction programming
func SumLoop(nums []int) int {
	// The sum variable is not immutable.
	// Immutable variable: A variable whose value cannot be
	// changed after it is assigned at runtime.
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
