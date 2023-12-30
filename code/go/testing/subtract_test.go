package matematica

import (
	"fmt"
	"testing"
)

func TestSubstract(t *testing.T) {
	if Subtract(4, 1) != 3 {
		t.Error("Failed to subtract 1 from 4")
	}
}

var TestsSubtraction = []struct {
	A, B   int
	Result int
}{
	{1, 2, -1},
	{2, 2, 0},
	{5, 2, 3},
}

func TestSubtractTable(t *testing.T) {
	for _, n := range TestsSubtraction {
		t.Run(fmt.Sprintf("%d-%d=%d", n.A, n.B, n.Result), func(t *testing.T) {
			if n.Result != Subtract(n.A, n.B) {
				t.Error("Failed")
			}
		})
	}
}
