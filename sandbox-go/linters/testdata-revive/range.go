// file-header-check
package testdata

func s(args []string) {
	for i, _ := range args {
		_ = i
		// _ = v
	}
}
