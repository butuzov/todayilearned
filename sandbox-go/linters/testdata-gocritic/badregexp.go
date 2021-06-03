package testdata_gocritic

import "regexp"

var re = regexp.MustCompile(`(?:^aa|bb|cc)foo[aba]`)
