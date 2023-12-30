// Test of empty-lines.

// file-header-check
package testdata

func f1(x *int) bool { // MATCH /extra empty line at the start of a block/

	return *x > 2
}
