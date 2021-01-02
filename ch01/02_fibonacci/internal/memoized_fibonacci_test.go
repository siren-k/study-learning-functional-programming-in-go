package internal

import "testing"

func TestMemoizedFibonacci(t *testing.T) {
	for _, ft := range FibonacciTests {
		if v := MemoizedFibonacci(ft.value); v != ft.expected {
			t.Errorf("MemoizedFibonacci(%d) return %d, expected %d", ft.value, v, ft.expected)
		}
	}
}

func BenchmarkMemoizedFibonacci(b *testing.B) {
	fn := MemoizedFibonacci
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkMemoizedFibonacci(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		MemoizedFibonacci(i)
	}
}

func BenchmarkMemoizedFibonacci0(b *testing.B)  { benchmarkMemoizedFibonacci(0, b) }
func BenchmarkMemoizedFibonacci1(b *testing.B)  { benchmarkMemoizedFibonacci(1, b) }
func BenchmarkMemoizedFibonacci2(b *testing.B)  { benchmarkMemoizedFibonacci(2, b) }
func BenchmarkMemoizedFibonacci3(b *testing.B)  { benchmarkMemoizedFibonacci(3, b) }
func BenchmarkMemoizedFibonacci4(b *testing.B)  { benchmarkMemoizedFibonacci(4, b) }
func BenchmarkMemoizedFibonacci5(b *testing.B)  { benchmarkMemoizedFibonacci(5, b) }
func BenchmarkMemoizedFibonacci6(b *testing.B)  { benchmarkMemoizedFibonacci(6, b) }
func BenchmarkMemoizedFibonacci10(b *testing.B) { benchmarkMemoizedFibonacci(10, b) }
func BenchmarkMemoizedFibonacci21(b *testing.B) { benchmarkMemoizedFibonacci(21, b) }
func BenchmarkMemoizedFibonacci43(b *testing.B) { benchmarkMemoizedFibonacci(43, b) }
