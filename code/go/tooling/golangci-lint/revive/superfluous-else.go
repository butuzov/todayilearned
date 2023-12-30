// Test of return+else warning.

// Package pkg ...
// file-header-check
package testdata

func h(f func() bool) string {
	for {
		if ok := f(); ok {
			a := 1
			_ = a
			continue
		} else { // MATCH /if block ends with a continue statement, so drop this else and outdent its block (move short variable declaration to its own line if necessary)/
			return "it's NOT okay!"
		}
	}
}
