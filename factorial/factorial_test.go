package factorial_test

import (
	. "github.com/daviskoh/go-code-challenges/factorial"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factorial", func() {
	Describe("Numbers less than 3", func() {
		It("should handle a 1", func() {
			Expect(Factorial(1)).To(Equal(1))
		})

		It("should handle 2", func() {
			Expect(Factorial(2)).To(Equal(2))
		})
	})

	Describe("All other numbers", func() {
		It("should handle 4", func() {
			Expect(Factorial(4)).To(Equal(4 * 3 * 2 * 1))
		})

		It("should handle 8", func() {
			Expect(Factorial(8)).To(Equal(8 * 7 * 6 * 5 * 4 * 3 * 2 * 1))
		})
	})
})
