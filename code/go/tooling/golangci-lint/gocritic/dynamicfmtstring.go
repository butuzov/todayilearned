package testdata_gocritic

import "fmt"

func fmtUndefinedFormatting() {
	foo := "foo happened"
	fooFunc := func(k string) (string, string) { return "", "123" }
	barFunc := func() string { return "123" }
	/*! use errors.New(foo) or fmt.Errorf("%s", foo) instead */
	fmt.Errorf(foo)

	/*! use errors.New(fooFunc("123")) or fmt.Errorf("%s", fooFunc("123")) instead */
	fmt.Errorf(fooFunc("123"))

	/*! use errors.New(barFunc()) or fmt.Errorf("%s", barFunc()) instead */
	fmt.Errorf(barFunc())
}
