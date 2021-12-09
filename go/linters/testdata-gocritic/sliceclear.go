package testdata_gocritic

func _(buf []int) {
	for i := 0; i < len(buf); i++ {
		buf[i] = 0
	}
}
