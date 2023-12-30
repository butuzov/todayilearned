// file-header-check
package testdata

func foo() (a, b, c, d int) { // MATCH /maximum number of return results per function exceeded; max 3 but got 4/
	// var a, b, c, d int
	return
}
