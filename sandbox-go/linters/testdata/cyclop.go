package testdata

import "strconv"

// FIXME: fails to be detected by linter

// PROJECT URL:
// https://github.com/bkielbasa/cyclop

func cyclop(n int) int {
	s := strconv.Itoa(n)
	if s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" {
		return n
	}
	if s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" {
		return n
	}
	if s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" {
		return n
	}

	return n
}
