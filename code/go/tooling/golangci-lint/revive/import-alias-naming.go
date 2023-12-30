// file-header-check
package testdata

import (
	. "strings"       // . aliases should be ignored
	_ "strings"       // _ aliases should be ignored
	bar_foo "strings" // MATCH /import name (bar_foo) must match the regular expression: ^[a-z][a-z0-9]{0,}$/
	fooBAR "strings"  // MATCH /import name (fooBAR) must match the regular expression: ^[a-z][a-z0-9]{0,}$/
	magical "strings"
	v1 "strings"
)

func somefunc() {
	fooBAR.Clone("")
	bar_foo.Clone("")
	v1.Clone("")
	magical.Clone("")
	Clone("")
}
