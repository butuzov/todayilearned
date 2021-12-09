package testdata_gocritic

// Example for the rule: unnecessaryBlock

func _(n int) {
	{
		print(n)
	}
}
