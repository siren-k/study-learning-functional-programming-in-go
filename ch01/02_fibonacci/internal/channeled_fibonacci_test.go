package internal

import "testing"

func TestChanneledFibonacci(t *testing.T) {
	for _, ft := range FibonacciTests {
		if v := ChanneledFibonacci(ft.value); v != ft.expected {
			t.Errorf("ChanneledFibonacci(%d) returned %d, expected %d", ft.value, v, ft.expected)
		}
	}
}

func BenchmarkChanneledFibonacci(b *testing.B) {
	fn := ChanneledFibonacci
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkChanneledFibonacci(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		ChanneledFibonacci(i)
	}
}

func BenchmarkChanneledFibonacci0(b *testing.B)  { benchmarkChanneledFibonacci(0, b) }
func BenchmarkChanneledFibonacci1(b *testing.B)  { benchmarkChanneledFibonacci(1, b) }
func BenchmarkChanneledFibonacci2(b *testing.B)  { benchmarkChanneledFibonacci(2, b) }
func BenchmarkChanneledFibonacci3(b *testing.B)  { benchmarkChanneledFibonacci(3, b) }
func BenchmarkChanneledFibonacci4(b *testing.B)  { benchmarkChanneledFibonacci(4, b) }
func BenchmarkChanneledFibonacci5(b *testing.B)  { benchmarkChanneledFibonacci(5, b) }
func BenchmarkChanneledFibonacci6(b *testing.B)  { benchmarkChanneledFibonacci(6, b) }
func BenchmarkChanneledFibonacci21(b *testing.B) { benchmarkChanneledFibonacci(21, b) }
func BenchmarkChanneledFibonacci43(b *testing.B) { benchmarkChanneledFibonacci(43, b) }
