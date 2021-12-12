package matematica

import (
	"fmt"
	"testing"
)

var TestCasesAdd = []struct {
	A, B   int
	Result int
}{
	{1, 2, 3},
	{2, 2, 4},
	{5, 2, 7},
}

func TestAddTable(t *testing.T) {

	for _, n := range TestCasesAdd {
		t.Run(fmt.Sprintf("%d+%d=%d", n.A, n.B, n.Result), func(t *testing.T) {
			t.Parallel()
			if n.Result != Add(n.A, n.B) {
				t.Fail()
			}
		})
	}
}

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range TestCasesAdd {
			Add(n.A, n.B)
		}
	}
}
