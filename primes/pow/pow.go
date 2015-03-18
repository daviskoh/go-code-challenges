package pow

func Pow(n1, n2 int) int {
	// if n2 < 2 then return n1
	if n2 < 2 {
		return n1
	}

	// return product of n1 * Pow(n1, n2 - 1)
	return n1 * Pow(n1, n2-1)
}
