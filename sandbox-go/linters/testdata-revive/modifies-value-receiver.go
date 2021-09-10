package testdata

type foozz struct {
	base string
}

func (f foozz) F() {
	f.base = "foo"
}
