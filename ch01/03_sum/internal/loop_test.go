package internal

import (
	"fmt"
	"testing"
)

func TestSumLoop(t *testing.T) {
	for _, tc := range TestCases {
		if v := SumLoop(tc.values); v != tc.expected {
			t.Errorf("SumLoop(%d) returned %d, expected %d", tc.values, v, tc.expected)
		}
	}
}

func ExampleSumLoop() {
	fmt.Println(SumLoop([]int{1}))
	fmt.Println(SumLoop([]int{1, 2}))
	fmt.Println(SumLoop([]int{1, 2, 3}))
	fmt.Println(SumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(SumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))
	fmt.Println(SumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}))
	// Output:
	// 1
	// 3
	// 6
	// 55
	// 210
	// 820
}

func BenchmarkSumLoop(b *testing.B) {
	fn := SumLoop
	for i := 0; i < b.N; i++ {
		_ = fn([]int{1, 2, 3})
	}
}

func benchmarkSumLoop(v []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumLoop(v)
	}
}

func BenchmarkSumLoop1(b *testing.B)  { benchmarkSumLoop([]int{1}, b) }
func BenchmarkSumLoop2(b *testing.B)  { benchmarkSumLoop([]int{1, 2}, b) }
func BenchmarkSumLoop3(b *testing.B)  { benchmarkSumLoop([]int{1, 2, 3}, b) }
func BenchmarkSumLoop4(b *testing.B)  { benchmarkSumLoop([]int{1, 2, 3, 4}, b) }
func BenchmarkSumLoop5(b *testing.B)  { benchmarkSumLoop([]int{1, 2, 3, 4, 5}, b) }
func BenchmarkSumLoop10(b *testing.B) { benchmarkSumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, b) }
func BenchmarkSumLoop20(b *testing.B) {
	benchmarkSumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, b)
}
func BenchmarkSumLoop40(b *testing.B) {
	benchmarkSumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, b)
}
