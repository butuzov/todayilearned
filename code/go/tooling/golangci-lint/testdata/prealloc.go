package testdata

// PROJECT URL: https://github.com/alexkohler/prealloc
func _() bool {
	existing := make([]int64, 10, 10)

	var init1 []int64
	for _, element := range existing {
		init1 = append(init1, element)
	}

	return len(init1) == len(existing)
}
