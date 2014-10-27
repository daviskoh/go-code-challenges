package prob2_test

import (
	. "github.com/daviskoh/go-code-challenges/project_euler/prob2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
var _ = Describe("FibonacciSumEven", func() {
	It("should handle base case 1", func() {
		Expect(FibonacciSum(1)).To(Equal(0))
	})

	It("should handle base case 2", func() {
		Expect(FibonacciSum(2)).To(Equal(2))
	})
})
