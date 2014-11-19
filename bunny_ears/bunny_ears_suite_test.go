package bunny_ears_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBunnyEars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BunnyEars Suite")
}
