package pig_latin_test

import (
	. "github.com/daviskoh/go-code-challenges/pig_latin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Vowels: a, e, i, o, u, y*
// y only vowel IF 1st letter

var _ = Describe("PigLatin", func() {
	It("should handle an empty string", func() {
		Expect(PigLatin("")).To(Equal(""))
	})

	It("should handle a single letter containing no vowels", func() {
		Expect(PigLatin("s")).To(Equal("say"))
	})

	It("should handle a multi letter string containing no vowels", func() {
		Expect(PigLatin("sty")).To(Equal("styay"))
	})

	It("should handle a single letter string starting with a vowel", func() {
		Expect(PigLatin("a")).To(Equal("away"))
	})

	It("should handle a multi letter string starting with a vowel", func() {
		Expect(PigLatin("emmm")).To(Equal("emmmway"))
	})

	It("should treat y as a vowel if it is 1st letter", func() {
		Expect(PigLatin("yzz")).To(Equal("yzzway"))
	})

	It("should handle string that has a non-first vowel occurrence", func() {
		Expect(PigLatin("story")).To(Equal("orystay"))
	})
})
