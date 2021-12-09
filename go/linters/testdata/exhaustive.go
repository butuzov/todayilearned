package testdata

// PROJECT URL: https://github.com/nishanths/exhaustive
type CompareResult int

const (
	Less CompareResult = iota - 1
	Equal
	Greater
	None
)

func (c CompareResult) String() string {
	switch c {
	case Less:
		return "Less"
	case Equal:
		return "Equal"
	case Greater:
		return "Greater"
	}

	return "None"
}
