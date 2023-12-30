// Test of redundant if err != nil

// Package pkg ...
// file-header-check
package testdata

func f() error {
	return nil
}

func g() error {
	if err := f(); err != nil { // MATCH /redundant if ...; err != nil check, just return error instead./
		return err
	}
	return nil
}
