// file-header-check
package testdata

//  MATCH /should not use the following blacklisted import: "crypto/md5"/
import (
	md5 "crypto/md5" // MATCH /Package "crypto/md5" already imported/
)

var _ = md5.Size
