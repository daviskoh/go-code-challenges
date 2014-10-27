package index_of_test

import (
	"fmt"

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
		Expect(IndexOf("a", "a", 0)).To(Equal(0))
	})

	It("should return -1 if string does not contain substring", func() {
		Expect(IndexOf("a", "b", 0)).To(Equal(-1))
	})

	It("should handle a string with char > 1", func() {
		Expect(IndexOf("ab", "b", 0)).To(Equal(1))
		Expect(IndexOf("aabc", "b", 0)).To(Equal(2))
	})

	It("should handle multi char substrings", func() {
		fmt.Println("***********************************")
		fmt.Println("MULTI")
		Expect(IndexOf("azbcasdf", "bc", 0)).To(Equal(2))

		fmt.Println("***********************************")
		fmt.Println("MULTI")
		Expect(IndexOf("zzzzzzz", "bc", 0)).To(Equal(-1))
	})
})
