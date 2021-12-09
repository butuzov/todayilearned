// file-header-check
package testdata

func foo() string {
	customId := "result"
	customVm := "result" // MATCH /var customVm should be customVM/

	invalid_name := 0

	return customId
}
