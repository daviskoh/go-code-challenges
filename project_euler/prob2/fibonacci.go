package prob2

func Fibonacci(sequencePos int) int {
	if sequencePos == 1 || 2 == sequencePos {
		return sequencePos
	}

	return Fibonacci(sequencePos-2) + Fibonacci(sequencePos-1)
}
