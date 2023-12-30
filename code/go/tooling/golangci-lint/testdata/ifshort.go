package testdata

import "github.com/pkg/errors"

// PROJECT URL: https://github.com/esimonov/ifshort
func ifshort(k string, m map[string]interface{}) error {
	_, ok := m[k]
	if !ok {
		return errors.New("not ok")
	}

	otherFunc1 := func() error {
		return errors.New("nope")
	}

	err := otherFunc1()
	if err != nil {
		return errors.Errorf("not ok: %w", err)
	}

	return nil
}

func DontUseShortSyntaxWhenPossible() {
	getValue := func() interface{} { return nil }

	v := getValue() // ERROR "variable 'v' is only used in the if-statement .*"
	if v != nil {
		return
	}
}
