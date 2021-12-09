package testdata

import "log"

// PROJECT URL: https://github.com/jirfag/go-printf-func-name

func _(format string, args ...interface{}) {
	const prefix = "[my] "

	log.Printf(prefix+format, args...)
}
