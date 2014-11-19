package bunny_ears_test

import (
	. "github.com/daviskoh/go-code-challenges/bunny_ears"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * We have a number of bunnies and each bunny has two big floppy ears.
 * We want to compute the total number of ears across all the bunnies recursively (without loops or multiplication).
 */

var _ = Describe("BunnyEars", func() {
	It("should handle no bunnies", func() {
		Expect(BunnyEars(0)).To(Equal(0))
	})

	It("should handle a single bunny", func() {
		Expect(BunnyEars(1)).To(Equal(2))
	})

	It("should handle 2 bunnies", func() {
		Expect(BunnyEars(2)).To(Equal(4))
	})

	It("should handle 4 bunnies", func() {
		Expect(BunnyEars(4)).To(Equal(8))
	})
})
