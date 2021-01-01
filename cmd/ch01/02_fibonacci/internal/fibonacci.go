package internal

// 파보나치 수열(Fibonacci Numbers)는 첫째 및 둘째 항이 1이며 그 뒤에 모든 항은 바로 앞 두 항의 합인 수열이다.
// 처음 여섯 항은 각각 1, 1, 2, 3, 5, 8이다. 편의상 0번째 항을 0으로 두기도 한다.
// 피보나치 수 Fn는 다음과 같은 초기값 및 점화식으로 정의된다.
// F₁ = F₂ = 1
// Fₙ = Fₙ₋₁ + Fₙ₋₂
// ==> 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, ...
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
