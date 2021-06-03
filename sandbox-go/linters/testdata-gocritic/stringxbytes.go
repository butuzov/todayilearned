package testdata_gocritic

// Example for the rule: stringXbytes

func _(s string) []byte {
	var b []byte
	copy(b, []byte(s))
	return b
}
