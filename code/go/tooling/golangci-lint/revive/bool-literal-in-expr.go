// file-header-check
package testdata

func foo(b, d int) bool {
	f := func() bool { return true }

	if (b == d) == true { // MATCH /omit Boolean literal in expression/
	}
	for f() || false != (b == d) { // MATCH /omit Boolean literal in expression/
	}

	return b > b == false // MATCH /omit Boolean literal in expression/
}

func isTrue(arg bool) bool {
	return arg
}

func main() {
	isTrue(true)
}
