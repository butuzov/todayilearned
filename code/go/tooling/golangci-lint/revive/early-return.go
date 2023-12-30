// Test of empty-blocks.

// file-header-check
package testdata

func earlyRet(cond bool) bool {
	if cond { //   MATCH /if c {...} else {... return } can be simplified to if !c { ... return } .../
		println()
		println()
		println()
	} else {
		return false
	}

	return true
}
