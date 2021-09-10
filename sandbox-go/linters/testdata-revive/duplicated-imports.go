package testdata

import (
	md5 "crypto/md5"    // MATCH /Package "crypto/md5" already imported/
	md5alt "crypto/md5" // MATCH /Package "crypto/md5" already imported/
)

func fooMd5(data []byte) bool {
	return md5.Sum(data) == md5alt.Sum(data)
}
