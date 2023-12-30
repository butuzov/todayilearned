package testdata

// PROJECT URL: https://github.com/ashanbrown/makezero
func _(nums []int) []int {
	values := make([]int, len(nums))
	for _, n := range nums {
		values = append(values, n)
	}

	return values
}
