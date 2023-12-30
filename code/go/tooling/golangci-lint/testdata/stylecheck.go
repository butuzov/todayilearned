package testdata

// PROJECT URL: https://github.com/dominikh/go-tools/tree/master/stylecheck
func Stylecheck(x int) {
	switch x {
	case 1:
		return
	default: // ERROR "ST1015: default case should be first or last in switch statement"
		return
	case 2:
		return
	}
}
