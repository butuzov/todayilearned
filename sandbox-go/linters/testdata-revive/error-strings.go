// file-header-check
package testdata

import "errors"

var errNew = errors.New("Foo asdasd asdad ad ad ads")

func rfooerr() error {
	return errors.New("Foo asdasd asdad ad ad ads")
}
