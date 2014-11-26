package fizzbuzz_test

import (
	. "github.com/daviskoh/go-code-challenges/fizzbuzz/fizzbuzz"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * func that checks value n against x & y:
 * if divisble by x then return "F"
 * if divisble by y then return "B"
 * if divisble by both then return "FB"
 */

var _ = Describe("Fizzbuzz", func() {
	It("handles non-matches", func() {
		Expect(FizzBuzz("2", "3", "5")).To(Equal("5"))
	})

	It("handles a simple num divisible by x", func() {
		Expect(FizzBuzz("2", "5", "4")).To(Equal("F"))
	})

	It("handles a simple num divisible by y", func() {
		Expect(FizzBuzz("9", "3", "6")).To(Equal("B"))
	})

	It("handles 0", func() {
		Expect(FizzBuzz("2", "3", "0")).To(Equal("0"))
	})
})
