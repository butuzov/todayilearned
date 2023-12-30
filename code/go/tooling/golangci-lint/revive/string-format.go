// Test string literal regex checks

// file-header-check
package testdata

func stringFormatMethod1(a, b string) {
}

func stringFormatMethod2(a, b string, c struct {
	d string
}) {
}

type stringFormatMethods struct{}

func (s stringFormatMethods) Method3(a, b, c string) {
}

func stringFormat() {
	stringFormatMethod1("This string is fine", "")
	stringFormatMethod1("this string is not capitalized", "") // MATCH /must start with a capital letter/
}
