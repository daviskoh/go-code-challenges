package prob1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProb1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prob1 Suite")
}
