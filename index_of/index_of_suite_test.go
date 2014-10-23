package index_of_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIndexOf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IndexOf Suite")
}
