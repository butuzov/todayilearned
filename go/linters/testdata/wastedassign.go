package testdata

// Project URL:
// https://github.com/sanposhiho/wastedassign

import (
	"strings"
)

func wastedassign(x int) int {
	return x + 1
}

func wastedassign_3(val interface{}, times uint) interface{} {
	switch hoge := val.(type) {
	case int:
		return 12
	case string:
		return strings.Repeat(hoge, int(times))
	default:
		return nil
	}
}

func wastedassign_2(params string) int {
	a := 12
	println(a)
	return a
}

func wastedassign_1(param int) int {
	println(param)
	useOutOfIf := 1212121 // ERROR "wasted assignment"
	ret := 0
	if false {
		useOutOfIf = 200 // ERROR "reassigned, but never used afterwards"
		return 0
	} else if param == 100 {
		useOutOfIf = 100 // ERROR "wasted assignment"
		useOutOfIf = 201
		useOutOfIf = wastedassign(useOutOfIf)
		useOutOfIf += 200 // ERROR "wasted assignment"
	} else {
		useOutOfIf = 100
		useOutOfIf += 100
		useOutOfIf = wastedassign(useOutOfIf)
		useOutOfIf += 200 // ERROR "wasted assignment"
	}

	if false {
		useOutOfIf = 200 // ERROR "reassigned, but never used afterwards"
		return 0
	} else if param == 200 {
		useOutOfIf = 100 // ERROR "wasted assignment"
		useOutOfIf = 201
		useOutOfIf = wastedassign(useOutOfIf)
		useOutOfIf += 200
	} else {
		useOutOfIf = 100
		useOutOfIf += 100
		useOutOfIf = wastedassign(useOutOfIf)
		useOutOfIf += 200
	}
	println(useOutOfIf)
	useOutOfIf = 192
	useOutOfIf += 100
	useOutOfIf += 200 // ERROR "reassigned, but never used afterwards"
	return ret
}

func wastedassign_4() int {
	hoge := 12
	noUse := 1111
	println(noUse)

	noUse = 1111 // ERROR "reassigned, but never used afterwards"
	for {
		if hoge == 14 {
			break
		}
		hoge = hoge + 1
	}
	return hoge
}

func wastedassign_6() {
	var i int
	var hoge int
	for {
		hoge = 5 // ERROR "reassigned, but never used afterwards"
	}

	println(i)
	println(hoge)
	return
}

func wastedassign_5() {
	var i int
	var hoge int
	for {
		hoge = 5
		break
	}

	println(i)
	println(hoge)
	return
}
