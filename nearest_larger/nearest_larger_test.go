package nearest_larger_test

import (
	. "github.com/daviskoh/go-code-challenges/nearest_larger"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Write a function, `nearest_larger(arr, i)` which takes an array and an
// index.  The function should return another index, `j`: this should
// satisfy:
//
// (a) `arr[i] < arr[j]`, AND
// (b) there is no `j2` closer to `i` than `j` where `arr[i] < arr[j]`.
//
// In case of ties (see example beow), choose the earliest (left-most)
// of the two indices. If no number in `arr` is largr than `arr[i]`,
// return `nil`.
//
// Difficulty: 2/5

var _ = Describe("NearestLarger", func() {
	It("should return -1 if no number in slice[i] is larger than slice[i]", func() {
		Expect(NearestLarger([]int{1, 1}, 0)).To(Equal(-1))
	})

	It("should return -1 if not other number in slice", func() {
		Expect(NearestLarger([]int{1}, 0)).To(Equal(-1))
	})

	It("should handle a simple case to the right", func() {
		Expect(NearestLarger([]int{1, 2}, 0)).To(Equal(1))
	})

	It("should handle a simple case to the left", func() {
		Expect(NearestLarger([]int{3, 2}, 1)).To(Equal(0))
	})

	It("should handle complex case to the right", func() {
		Expect(NearestLarger([]int{2, 3, 4, 4, 5}, 2)).To(Equal(4))
	})

	It("should handle complex case to the left", func() {
		Expect(NearestLarger([]int{2, 7, 2, 4, 3}, 3)).To(Equal(1))
	})

	It("upon a tie, it should choose left index", func() {
		Expect(NearestLarger([]int{2, 4, 3, 4, 7}, 2)).To(Equal(1))
	})
})
