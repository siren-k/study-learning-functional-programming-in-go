package internal

func channel(ch chan int, counter int) {
	n1, n2 := 0, 1
	for i := 0; i < counter; i++ {
		ch <- n1
		n1, n2 = n2, n1+n2
	}
	close(ch)
}

// The next number calculation process in the Fibonacci sequence is a small factor
// that can be executed in parallel. Channels are a good example of concurrency
// implementation using Fibonacci sequence calculations.
func ChanneledFibonacci(n int) int {
	n += 1
	// Go's channel works like a first-in, first-out(FIFO) queue.
	ch := make(chan int)

	// Parallelism: running multiple functions simultaneously on different CPU cores
	// Concurrency: Dividing the program into independently executable operations
	go channel(ch, n)
	i := 0
	var result int
	for num := range ch {
		result = num
		i++
	}
	return result
}
