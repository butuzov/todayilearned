// file-header-check
package testdata

func fooVarNaming() string {
	customId := "result"
	customVm := "result" // MATCH /var customVm should be customVM/
	_ = customVm
	invalid_name := 0
	_ = invalid_name
	return customId
}
