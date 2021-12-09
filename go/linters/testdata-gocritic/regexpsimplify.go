package testdata_gocritic

import "regexp"

var reSimplify = regexp.MustCompile(`(?:a|b|c)   [a-z][a-z]*`)
