// file-header-check
package testdata

// from github.com/ugorji/go/codec/helper.go
func isNaN(f float64) bool { return f != f } // MATCH /expression always evaluates to false/
