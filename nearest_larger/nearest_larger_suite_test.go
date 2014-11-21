package nearest_larger_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNearestLarger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NearestLarger Suite")
}
