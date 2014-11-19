package factorial

func Factorial(n int) int {
	if n < 3 {
		return n
	}

	return n * Factorial(n-1)
}
