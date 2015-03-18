package primes

import (
	"testing"
)

func compareSlices(t *testing.T, slice1, slice2 []int) {
	slice1Length := len(slice1)
	slice2Length := len(slice2)

	greater := slice2Length

	if slice1Length > slice2Length {
		greater = slice1Length
	}

	for i := 0; i < greater; i++ {
		if slice1[i] != slice2[i] {
			defer t.Errorf("slices mismatched at index: %i with values %i & %i", i, slice1[i], slice2[i])
		}
	}
}

func TestRejectNegativeIntegers(t *testing.T) {
	negativeIntegers := []int{
		-4,
		-3,
		-2,
		-2,
	}

	expected := []int{}

	for _, negInt := range negativeIntegers {
		compareSlices(t, Primes(negInt), expected)
	}
}

func TestMaxNumbersThatDoNotHavePrimes(t *testing.T) {
	nums := []int{
		1,
		2,
	}

	expected := []int{}

	for _, num := range nums {
		compareSlices(t, Primes(num), expected)
	}
}

func TestNumbersThatHavePrimes(t *testing.T) {
	compareSlices(t, Primes(3), []int{2})

	compareSlices(t, Primes(4), []int{3, 2})

	compareSlices(t, Primes(6), []int{5, 3, 2})
}
