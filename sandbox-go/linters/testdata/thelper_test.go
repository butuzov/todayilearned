package testdata

import "testing"

// PROJECT URL: https://github.com/kulti/thelper

func TestBad(t *testing.T) {
	testBad(t)
}

func testBad(t *testing.T) {
	t.Fatal("unhelpful line") // <--- output point this line
}
