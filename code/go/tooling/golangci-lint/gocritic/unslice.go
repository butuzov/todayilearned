package testdata_gocritic

// Example for the rule: unslice

func _(s string, b []byte, v ...interface{}) {
	f7(s[:]) // s is string
	// this is triggering error by golang.
	// copy(b[:], v...) // b is []byte
}

func f7(s string) {}
