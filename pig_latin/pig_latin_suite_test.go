package pig_latin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPigLatin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PigLatin Suite")
}
