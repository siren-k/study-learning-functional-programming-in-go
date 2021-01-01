package internal

// Pure functional programming
func SumRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// Pure functional programming languages traverse enum elements recursively.
	return nums[0] + SumRecursive(nums[1:])
}
