package isprime

/**
 * Wikipedia:
 * A prime number (or a prime) is a natural number
 * greater than 1 that has no positive divisors
 * other than 1 and itself.
 */

func IsPrime(n int) bool {
	// divide n by all numbers
	// from itself - 1 down through 2

	if n < 2 {
		return false
	}

	for i := n - 1; i >= 2; i-- {
		if n%i == 0 {
			return false
		}
	}

	return true
}
