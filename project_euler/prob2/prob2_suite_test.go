package prob2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProb2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prob2 Suite")
}
