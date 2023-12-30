// file-header-check
package testdata

func conditionalExpr(x, y bool) bool {
	_ = x == y  // should not match, not AND or OR operators
	if x || y { // should not match, no caller
		return true
	}
	foo := caller(x, y) || y // MATCH /for better performance 'caller(x, y) || y' might be rewritten as 'y || caller(x, y)'/
	if caller(x, y) || y {   // MATCH /for better performance 'caller(x, y) || y' might be rewritten as 'y || caller(x, y)'/
		return true
	}

	return foo
}

func caller(x, y bool) bool {
	return true
}
