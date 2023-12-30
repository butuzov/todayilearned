package testdata

var (
	decoa  = 1
	deco_c = 1
	decob  = 1 // want "multiple \"var\" declarations are not allowed; use parentheses instead"
)

func decof() {
	const decog = 1
}

type decoe int // want "type must not be placed after const"

func init() {} // want "init func must be the first function in file"
