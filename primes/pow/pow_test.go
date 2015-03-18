package pow

import (
	"../testingutils"
	"testing"
)

func TestNumberRaisedToItself(t *testing.T) {
	var n int

	n = Pow(1, 1)
	testingutils.Expect(t, n, 1*1)

	n = Pow(2, 2)
	testingutils.Expect(t, n, 2*2)
}

func TestNumberRaisedToAnotherNum(t *testing.T) {
	n := Pow(2, 3)
	testingutils.Expect(t, n, 2*2*2)
}
