package internal

import (
	"fmt"
	"testing"
)

var TestCasesForCurryAddTwo = []struct {
	value  int
	expect int
}{
	{1, 3},
	{2, 4},
	{3, 5},
}

var TestCasesForCurryAddThree = []struct {
	value  int
	expect int
}{
	{1, 4},
	{2, 5},
	{3, 6},
}

func TestCurryAddTwo(t *testing.T) {
	fn := CurryAddTwo
	for _, tc := range TestCasesForCurryAddTwo {
		if v := fn(tc.value); v != tc.expect {
			t.Errorf("CurryAddTwo(%d) returned %d, expected %d", tc.value, v, tc.expect)
		}
	}
}

func ExampleCurryAddTwo() {
	fmt.Println(CurryAddTwo(1))
	fmt.Println(CurryAddTwo(2))
	fmt.Println(CurryAddTwo(3))

	// Output:
	// 3
	// 4
	// 5
}

func BenchmarkCurryAddTwo(b *testing.B) {
	fn := CurryAddTwo
	for i := 0; i < b.N; i++ {
		_ = fn(1)
	}
}

func benchmarkCurryAddTwo(i int, b *testing.B) {
	fn := CurryAddTwo
	for n := 0; n < b.N; n++ {
		fn(i)
	}
}

func BenchmarkCurryAddTwo0(b *testing.B) { benchmarkCurryAddTwo(0, b) }
func BenchmarkCurryAddTwo1(b *testing.B) { benchmarkCurryAddTwo(1, b) }
func BenchmarkCurryAddTwo2(b *testing.B) { benchmarkCurryAddTwo(2, b) }
func BenchmarkCurryAddTwo3(b *testing.B) { benchmarkCurryAddTwo(3, b) }

func TestCurryAddThree(t *testing.T) {
	fn := CurryAddThree
	for _, tc := range TestCasesForCurryAddThree {
		if v := fn(tc.value); v != tc.expect {
			t.Errorf("CurryAddThree(%d) returned %d, expected %d", tc.value, v, tc.expect)
		}
	}
}

func ExampleCurryAddThree() {
	fmt.Println(CurryAddThree(1))
	fmt.Println(CurryAddThree(2))
	fmt.Println(CurryAddThree(3))

	// Output:
	// 4
	// 5
	// 6
}

func BenchmarkCurryAddThree(b *testing.B) {
	fn := CurryAddThree
	for n := 0; n < b.N; n++ {
		_ = fn(1)
	}
}

func benchmarkCurryAddThree(i int, b *testing.B) {
	fn := CurryAddThree
	for n := 0; n < b.N; n++ {
		fn(i)
	}
}

func BenchmarkCurryAddThree0(b *testing.B) { benchmarkCurryAddThree(0, b) }
func BenchmarkCurryAddThree1(b *testing.B) { benchmarkCurryAddThree(1, b) }
func BenchmarkCurryAddThree2(b *testing.B) { benchmarkCurryAddThree(2, b) }
func BenchmarkCurryAddThree3(b *testing.B) { benchmarkCurryAddThree(3, b) }
