package primes

import (
	// NO RELATIVE IMPORT????
	// "./isprime"
	"github.com/daviskoh/go-code-challenges/primes/isprime"
)

func Primes(n int) []int {
	primes := []int{}

	for i := n - 1; i >= 2; i-- {
		if isprime.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}
