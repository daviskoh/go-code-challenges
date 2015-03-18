package isprime

import (
	"../testingutils"
	"testing"
)

func TestInvalidNumbers(t *testing.T) {
	invalidNums := []int{
		-2,
		-1,
		0,
	}

	for _, invalidNum := range invalidNums {
		testingutils.Expect(t, IsPrime(invalidNum), false)
	}
}

func TestNumberTwo(t *testing.T) {
	testingutils.Expect(t, IsPrime(2), true)
}

func TestPrimeNumbers(t *testing.T) {
	primeNums := []int{
		3,
		5,
		7,
		11,
		13,
		17,
		19,
		23,
		29,
		31,
		37,
		41,
		43,
		47,
		53,
		59,
		61,
		67,
		71,
	}

	for _, primeNum := range primeNums {
		testingutils.Expect(t, IsPrime(primeNum), true)
	}
}

func TestNonPrimeNumbers(t *testing.T) {
	nonPrimes := []int{
		1,
		4,
		6,
		8,
		10,
		12,
	}

	for _, nonPrimeNum := range nonPrimes {
		testingutils.Expect(t, IsPrime(nonPrimeNum), false)
	}
}
