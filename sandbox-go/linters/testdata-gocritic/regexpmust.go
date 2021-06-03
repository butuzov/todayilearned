package testdata_gocritic

import "regexp"

// Example for the rule: regexpMust

func _(n int) {
	re, _ := regexp.Compile("const pattern")
	re.FindAllString("sdosdos", -1)
}
