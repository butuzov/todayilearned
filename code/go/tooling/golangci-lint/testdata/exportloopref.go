package testdata

// PROJECT URL: https://github.com/kyoh86/exportloopref
func _() {
	values := []string{"a", "b", "c"}
	var copies []*string
	for _, val := range values {
		copies = append(copies, &val)
	}
}
