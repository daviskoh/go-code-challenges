package example_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExample(t *testing.T) {
	/**
	 * A Ginkgo test signals failure by calling Ginkgo’s Fail(description string) function.
	 * We pass this function to Gomega using RegisterFailHandler.
	 * This is the sole connection point between Ginkgo and Gomeg
	 */
	RegisterFailHandler(Fail)

	/**
	 * tells Ginkgo to start the test suite. Ginkgo will automatically fail the testing.T if any of your specs fail.
	 */
	RunSpecs(t, "Example Suite")
}
