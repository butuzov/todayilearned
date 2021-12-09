package testdata_gocritic

// Example for the rule: mapKey

var out = map[string]int{
	" foo": 1,
	"foo":  2,
	"bar":  3,
	"bar ": 4,
}
