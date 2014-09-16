package example_test

import (
	. "github.com/daviskoh/go-code-challenges/example"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example", func() {
	It("does stuff", func() {
		Expect(DoStuff()).To(Equal(1))
	})
})
