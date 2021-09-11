// file-header-check
package testdata

func bar() {
	a := 1

	switch a {
	case 1, 2:
		a++
	}

loop:
	for {
		switch a {
		case 1:
			a++
			println("one")
			break // MATCH /omit unnecessary break at the end of case clause/
		case 2:
			println("two")
			break loop
		default:
			println("default")
		}
	}

	return // MATCH /omit unnecessary return statement/
}
