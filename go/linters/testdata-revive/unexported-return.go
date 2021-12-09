// file-header-check
package testdata

type foobarfoo struct{}

//revive:disable:unexported-return
func NewFooBar() *foobarfoo {
	return new(foobarfoo)
}

//revive:enable:unexported-return
