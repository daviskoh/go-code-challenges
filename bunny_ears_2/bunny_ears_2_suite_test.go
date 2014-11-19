package bunny_ears_2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBunnyEars2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BunnyEars2 Suite")
}
