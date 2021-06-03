package testdata_gocritic

// Example for the rule: nestingReduce

type Typos struct {
	Bool bool
}

func body()

func _(a []Typos) {
	for _, v := range a {
		if v.Bool {
			body()
		}
	}
}
