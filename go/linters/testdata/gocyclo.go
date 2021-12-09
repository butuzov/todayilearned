package testdata

// PROJECT URL: https://github.com/alecthomas/gocyclo
func _(n int) int {
	if n == 1 && n < 10 {
		if n > 10 {
			n += 10
		} else {
			n -= 10
		}
	} else if n > 1 {
		n *= 10
	} else if n > 10 {
		n *= 100
	} else if n > 1100 {
		n *= 110
	} else if n > 2100 {
		n *= 110
	} else if n > 3100 {
		n *= 110
	} else if n > 4100 {
		n *= 110
	} else if n > 5100 {
		n *= 110
	} else if n > 9100 {
		n *= 110
	}
	return n
}
