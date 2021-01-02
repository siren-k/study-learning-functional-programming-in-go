package internal

var TestCases = []struct {
	values   []int
	expected int
}{
	{[]int{1}, 1},
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3}, 6},
	{[]int{1, 2, 3, 4}, 10},
	{[]int{1, 2, 3, 4, 5}, 15},
}
