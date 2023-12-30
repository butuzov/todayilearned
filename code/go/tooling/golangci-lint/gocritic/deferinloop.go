package testdata_gocritic

func deferWithCall() {
	for {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer println("test")
		break
	}

	for range []int{1, 2, 3, 4} {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer println("test")
	}
}
