package factorial_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFactorial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factorial Suite")
}
