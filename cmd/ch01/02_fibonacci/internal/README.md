# _Test run and results_

---

> Unit tests should be written for the following purposes.
>> * Ensure that the implementation meets the functional requirements
>> * Use of testing to devise the best solution when implementing a solution
>> * Generate quality tests for use in continuous integration processes
>> * Verification that the implementation meets the application's interface requirements
>> * Easy integration test development
>> * Protect the code you are developing as the code you are developing may malfunction 
     due to code from other developers.

---

```
❯ go test -bench=. *.go -benchmem
goos: darwin
goarch: amd64
BenchmarkChanneledFibonacci-12            502346              2191 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci0-12          2290202               501 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci1-12          1361276               873 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci2-12          1291795               931 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci3-12           861127              1312 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci4-12           906406              1337 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci5-12           655208              1697 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci6-12           608020              1708 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci21-12          232143              4876 ns/op              96 B/op          1 allocs/op
BenchmarkChanneledFibonacci43-12          128187              8983 ns/op              96 B/op          1 allocs/op      ==> 채널 사용 버전: 8,983 ns
BenchmarkFibonacci-12                   18772230                61.9 ns/op             0 B/op          0 allocs/op
BenchmarkFibonacci0-12                  788530491                1.69 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci1-12                  704881917                1.64 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci2-12                  729197100                1.65 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci3-12                  276205075                4.16 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci4-12                  158625286                7.50 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci5-12                  83460087                13.7 ns/op             0 B/op          0 allocs/op
BenchmarkFibonacci6-12                  53914550                22.4 ns/op             0 B/op          0 allocs/op
BenchmarkFibonacci21-12                    35036             34909 ns/op               0 B/op          0 allocs/op
BenchmarkFibonacci43-12                        1        1361411132 ns/op               0 B/op          0 allocs/op      ==> 재귀함수 버전: 1,361,411,132 ns
BenchmarkMemoizedFibonacci-12           88230966                12.7 ns/op             0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci0-12          139185478                8.49 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci1-12          140355387                8.29 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci2-12          134812435                8.78 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci3-12          135455822                8.78 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci4-12          125055921                9.30 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci5-12          124786507                9.56 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci6-12          100000000               10.0 ns/op             0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci10-12         125661056                9.57 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci21-12         129347080                9.21 ns/op            0 B/op          0 allocs/op
BenchmarkMemoizedFibonacci43-12         100000000               10.2 ns/op             0 B/op          0 allocs/op      ==> 메모이즈 사용 버전: 10.2 ns
PASS
ok      command-line-arguments  52.175s
```
