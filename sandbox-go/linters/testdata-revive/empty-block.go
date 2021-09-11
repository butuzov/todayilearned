// file-header-check
package testdata

func g(f func() bool) {
	{ // MATCH /this block is empty, you can remove it/
	}

	// issue 386, then overwritten by issue 416
	c := make(chan int)
	for range c { // MATCH /this block is empty, you can remove it/
	}

	s := "a string"
	for range s { // MATCH /this block is empty, you can remove it/
	}
}
