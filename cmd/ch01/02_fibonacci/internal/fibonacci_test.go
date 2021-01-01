package internal

import "testing"

var FibonacciTests = []struct {
	value    int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{21, 10946},
	{43, 433494437},
}

func TestFibonacci(t *testing.T) {
	for _, ft := range FibonacciTests {
		if got := Fibonacci(ft.value); got != ft.expected {
			t.Errorf("Fibonacci(%d) returned %d, expected %d", ft.value, got, ft.expected)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	fn := Fibonacci
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibonacci(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(i)
	}
}

func BenchmarkFibonacci0(b *testing.B)  { benchmarkFibonacci(0, b) }
func BenchmarkFibonacci1(b *testing.B)  { benchmarkFibonacci(1, b) }
func BenchmarkFibonacci2(b *testing.B)  { benchmarkFibonacci(2, b) }
func BenchmarkFibonacci3(b *testing.B)  { benchmarkFibonacci(3, b) }
func BenchmarkFibonacci4(b *testing.B)  { benchmarkFibonacci(4, b) }
func BenchmarkFibonacci5(b *testing.B)  { benchmarkFibonacci(5, b) }
func BenchmarkFibonacci6(b *testing.B)  { benchmarkFibonacci(6, b) }
func BenchmarkFibonacci21(b *testing.B) { benchmarkFibonacci(21, b) }
func BenchmarkFibonacci43(b *testing.B) { benchmarkFibonacci(43, b) }
