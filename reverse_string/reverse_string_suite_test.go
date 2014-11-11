package reverse_string_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestReverseString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReverseString Suite")
}
