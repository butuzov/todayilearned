package testdata

// PROJECT URL: https://github.com/mvdan/unparam
func ClosureUse() {
	var enclosed int
	setValue := func(v *int) {
		enclosed = 2
	}
	var newValue int = 4
	println(enclosed)
	setValue(&newValue)
	println(enclosed)
}
