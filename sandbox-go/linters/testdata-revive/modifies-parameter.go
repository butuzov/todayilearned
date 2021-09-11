// file-header-check
package testdata

const (
	zero = 0.0
	ones = 1.1
)

func two(b, c float32) {
	if c > zero {
		b = ones * 2
	}
}
