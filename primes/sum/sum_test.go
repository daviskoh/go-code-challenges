package sum

import (
	"../testingutils"
	"testing"
)

func TestEmptySlice(t *testing.T) {
	sum := Sum([]int{})

	testingutils.Expect(t, sum, 0)
}

func TestSliceWithSingleValue(t *testing.T) {
	sum := Sum([]int{5})

	testingutils.Expect(t, sum, 5)
}

func TestMultipleValues(t *testing.T) {
	sum := Sum([]int{1, 2, 3, 4, 5, 6, 10})

	testingutils.Expect(t, sum, 1+2+3+4+5+6+10)
}
