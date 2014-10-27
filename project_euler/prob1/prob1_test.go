package prob1_test

import (
	. "github.com/daviskoh/go-code-challenges/project_euler/prob1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
 * The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

var _ = Describe("Prob1", func() {
	It("handles a simple case w/ 1 multiple of 3", func() {
		Expect(Prob1(3)).To(Equal(0))
	})

	It("handles a simple case w/ 1 multiple of 5", func() {
		Expect(Prob1(5)).To(Equal(3))
	})

	It("handles a value with 1 multiple of 3 & 1 multiple of 5", func() {
		Expect(Prob1(6)).To(Equal(3 + 5))
	})

	It("handles example from question", func() {
		Expect(Prob1(10)).To(Equal(23))
	})

	It("should get answer to question", func() {
		Expect(Prob1(1000)).To(Equal(233168))
	})
})
