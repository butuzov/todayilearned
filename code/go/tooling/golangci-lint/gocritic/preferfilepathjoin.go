package testdata_gocritic

import "os"

var (
	x  = "foo"
	y  = "bar"
	v1 = x + string(os.PathSeparator) + y
)
