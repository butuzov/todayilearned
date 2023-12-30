// file-header-check
package testdata

import (
	"time"
)

func ext۰time۰Sleep(fr string, args []any) { // MATCH /parameter 'fr' seems to be unused, consider removing or renaming it as _/
	time.Sleep(time.Duration(args[0].(int64)))
}
