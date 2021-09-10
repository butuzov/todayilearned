package testdata

//  MATCH /should not use the following blacklisted import: "crypto/md5"/
import (
	md5 "crypto/md5"    // MATCH /Package "crypto/md5" already imported/
	"crypto/md5" // MATCH /Package "crypto/md5" already imported/
	md5alt "crypto/md5" // MATCH /Package "crypto/md5" already imported/
)