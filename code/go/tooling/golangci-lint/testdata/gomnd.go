package testdata

// PROJECT URL:  https://github.com/tommy-muehle/go-mnd
func _(x int) bool {
	x = x * 100

	if x > 7 {
		return false
	}

	return x > 100
}
