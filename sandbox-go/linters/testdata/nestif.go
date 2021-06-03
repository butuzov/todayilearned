package testdata

// PROJECT URL:  https://github.com/nakabonne/nestif

func _() {
	var (
		baz       string
		qux, quux bool
		x         int
	)
	if baz == "baz" {
		if qux {
			if quux {
				if qux {
					x = 1
				}
			}
		}
	}
	print(x)
}
