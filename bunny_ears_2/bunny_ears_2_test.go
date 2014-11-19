package bunny_ears_2_test

import (
	. "github.com/daviskoh/go-code-challenges/bunny_ears_2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// We have bunnies standing in a line, numbered 1, 2, ... The odd bunnies (1, 3, ..) have the normal 2 ears.
// The even bunnies (2, 4, ..) we'll say have 3 ears, because they each have a raised foot.
// Recursively return the number of "ears" in the bunny line 1, 2, ... n (without loops or multiplication).

var _ = Describe("BunnyEars2", func() {
	It("should handle 0 bunnies", func() {
		Expect(BunnyEars2(0)).To(Equal(0))
	})

	It("should handle a single odd numbered bunny", func() {
		Expect(BunnyEars2(1)).To(Equal(2))
	})

	It("should handle a single odd & single even", func() {
		Expect(BunnyEars2(2)).To(Equal(2 + 3))
	})
})
