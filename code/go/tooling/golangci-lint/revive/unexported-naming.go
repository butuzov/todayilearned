// file-header-check
package testdata

var (
	unexported string
	Exported   string
)

func unexportednaming(
	S int, // MATCH /the symbol S is local, its name should start with a lowercase letter/
	s int,
) (
	Result bool, // MATCH /the symbol Result is local, its name should start with a lowercase letter/
	result bool,
) {
	const NotExportableConstant = "local" // MATCH /the symbol NotExportableConstant is local, its name should start with a lowercase letter/

	return S == 1, s == 2
}
