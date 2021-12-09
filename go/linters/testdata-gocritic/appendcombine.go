package testdata_gocritic

// Example for the rule: appendCombine

func _(n int) {
	var nums []int

	nums = append(nums, n)
	nums = append(nums, n)

	_ = nums
}
