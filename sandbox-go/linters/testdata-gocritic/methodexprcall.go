package testdata_gocritic

// Example for the rule: methodExprCall

type foobar struct{}

func (f1 foobar) bar() {}

func _() {
	f := foobar{}
	foobar.bar(f)
}
