package testdata_gocritic

import "reflect"

// Example for the rule: emptyFallthrough

func _(n reflect.Kind) reflect.Kind {
	switch n {
	case reflect.Int:
		fallthrough
	case reflect.Int32:
		return reflect.Int
	}

	return reflect.Int
}
