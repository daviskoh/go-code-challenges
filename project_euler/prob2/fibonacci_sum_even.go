package prob2

func FibonacciSum(sequencePos int) int {
	if sequencePos < 2 {
		return 0
	}

	fiboNum := Fibonacci(sequencePos)

	if fiboNum%2 == 0 {
		return fiboNum + FibonacciSum(sequencePos-1)
	}

	return FibonacciSum(sequencePos - 1)
}

// until fibo value > 4m
// keep adding fibo nums
