// file-header-check
package testdata

import "time"

var (
	fiveMins = 5 * time.Minute // skips internal
	TenHours = 10 * time.Minute * 60
	TenSec   = 10 * time.Second * 60
)
