package testdata_gocritic

import "sort"

var keys []int

func _(xs []int) {
	sort.Slice(xs, func(i, j int) bool { return keys[i] < keys[j] })
}
