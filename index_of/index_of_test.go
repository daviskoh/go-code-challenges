package index_of_test

import (
	. "github.com/daviskoh/go-code-challenges/index_of"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * Implement a function that takes a string, substring and outputs index of that substring
 * f(string, string) -> int
 */

var _ = Describe("IndexOf", func() {
	It("should handle a simple case with a 1 char string", func() {
		Expect(IndexOf("a", "a")).To(Equal(0))
	})

	It("should return -1 if string does not contain substring", func() {
		Expect(IndexOf("a", "b")).To(Equal(-1))
	})
})
