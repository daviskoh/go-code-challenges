package testingutils

import (
	"testing"
)

func Expect(t *testing.T, pow, n interface{}) {
	if pow != n {
		t.Errorf("Expect %i = %i", pow, n)
	}
}
