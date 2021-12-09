package testdata_gocritic

import "strings"

var (
	x1, y1 string
	_      = strings.Join([]string{x1, y1}, "_")
)
