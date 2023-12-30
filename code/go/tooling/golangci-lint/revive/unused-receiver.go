// file-header-check
package testdata

type Unix struct{}

func (f *Unix) Name() string { // MATCH /method receiver 'f' is not referenced in method's body, consider removing or renaming it as _/
	return "unix"
}

func (u *Unix) Foo() (string, error) { // MATCH /method receiver 'u' is not referenced in method's body, consider removing or renaming it as _/
	{
		u := 1
		_ = u
	}
	return "", nil
}
