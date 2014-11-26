package parser_test

import (
	"github.com/daviskoh/go-code-challenges/fizzbuzz/fizzbuzz"
	. "github.com/daviskoh/go-code-challenges/fizzbuzz/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var formatter = fizzbuzz.FizzBuzz

	It("should handle a simple single num case", func() {
		Expect(Parser("2 3 1", formatter)).To(Equal("1"))
	})

	It("handles a multi num case", func() {
		Expect(Parser("2 3 4", formatter)).To(Equal("1 F B F"))
	})
})
