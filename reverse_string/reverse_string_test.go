package reverse_string_test

import (
	. "github.com/daviskoh/go-code-challenges/reverse_string"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReverseString", func() {
	It("should handle an empty string", func() {
		Expect(ReverseString("")).To(Equal(""))
	})

	It("should handle a 1 letter string", func() {
		Expect(ReverseString("a")).To(Equal("a"))
		Expect(ReverseString("b")).To(Equal("b"))
	})

	It("should handle 2 letter string", func() {
		Expect(ReverseString("ab")).To(Equal("ba"))
	})

	It("should handle multi letter strings", func() {
		Expect(ReverseString("abc123")).To(Equal("321cba"))
	})

	It("should handle non-alphanumeric chars", func() {
		Expect(ReverseString("!@# $")).To(Equal("$ #@!"))
	})
})
