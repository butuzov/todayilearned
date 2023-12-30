// file-header-check
package testdata

var foo any = "foo"

func handleIgnored() {
	_, _ = foo.(int) // MATCH /type cast result is unchecked in foo.(int) - type assertion result ignored/
}
