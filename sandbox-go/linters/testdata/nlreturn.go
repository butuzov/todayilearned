package testdata

// PROJECT URL:  https://github.com/ssgreg/nlreturn
func foo() int {
	a := 0
	_ = a
	return a
}

func bar() int {
	a := 0
	if a == 0 {
		_ = a
		return a
	}

	return a
}
