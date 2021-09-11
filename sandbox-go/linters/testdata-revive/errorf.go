// file-header-check
package testdata

import (
	"errors"
	"fmt"
)

var errOpps = errors.New(fmt.Sprintf("s: %s", "s"))
