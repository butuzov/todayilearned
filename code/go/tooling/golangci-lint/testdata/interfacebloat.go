package testdata

type InterfaceBloatExample01 interface { // want "the interface has more than 10 methods: 11"
	a01()
	a02()
	a03()
	a04()
	a05()
	a06()
	a07()
	a08()
	a09()
	a10()
	a11()
}
