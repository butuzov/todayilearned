package testdata_gocritic

import (
	"sort"
	"strings"
)

func _(s string) {
	xs := strings.Split(s, " ")
	xs = sort.StringSlice(xs)
}
