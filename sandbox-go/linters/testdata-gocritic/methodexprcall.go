package testdata_gocritic

// Example for the rule: methodExprCall

type foo struct{}

func (f1 foo) bar(f2 foo) {}

func _() {
	// f := foo{}
	// foo.bar(f)
}
