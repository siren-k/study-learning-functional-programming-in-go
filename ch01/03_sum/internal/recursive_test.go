package internal

import (
	"fmt"
	"testing"
)

func TestSumRecursive(t *testing.T) {
	for _, tc := range TestCases {
		if v := SumRecursive(tc.values); v != tc.expected {
			t.Errorf("SumRecursive(%d) returned %d, expected %d", tc.values, v, tc.expected)
		}
	}
}

func ExampleSumRecursive() {
	fmt.Println(SumRecursive([]int{1}))
	fmt.Println(SumRecursive([]int{1, 2}))
	fmt.Println(SumRecursive([]int{1, 2, 3}))
	fmt.Println(SumRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(SumRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))
	fmt.Println(SumRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}))
	// Output:
	// 1
	// 3
	// 6
	// 55
	// 210
	// 820
}

func BenchmarkSumRecursive(b *testing.B) {
	fn := SumRecursive
	for i := 0; i < b.N; i++ {
		_ = fn([]int{1, 2, 3})
	}
}

func benchmarkSumRecursive(v []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumRecursive(v)
	}
}

func BenchmarkSumRecursive1(b *testing.B) { benchmarkSumRecursive([]int{1}, b) }
func BenchmarkSumRecursive2(b *testing.B) { benchmarkSumRecursive([]int{1, 2}, b) }
func BenchmarkSumRecursive3(b *testing.B) { benchmarkSumRecursive([]int{1, 2, 3}, b) }
func BenchmarkSumRecursive4(b *testing.B) { benchmarkSumRecursive([]int{1, 2, 3, 4}, b) }
func BenchmarkSumRecursive5(b *testing.B) { benchmarkSumRecursive([]int{1, 2, 3, 4, 5}, b) }
func BenchmarkSumRecursive10(b *testing.B) {
	benchmarkSumRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, b)
}
func BenchmarkSumRecursive20(b *testing.B) {
	benchmarkSumRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, b)
}
func BenchmarkSumRecursive40(b *testing.B) {
	benchmarkSumRecursive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, b)
}
