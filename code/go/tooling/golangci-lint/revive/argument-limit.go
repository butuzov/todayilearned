// file-header-check
package testdata

func foobar_arglim(a, b, c, d int) { // MATCH /maximum number of arguments per function exceeded; max 3 but got 4/
}

func bar_arglim(a, b int) {
}

func baz_arglim(a string, b int) {
}

func qux(a string, b int, c int, d string, e int64) { // MATCH /maximum number of arguments per function exceeded; max 3 but got 5/
}
