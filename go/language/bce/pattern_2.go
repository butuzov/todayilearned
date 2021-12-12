package main

func p2(b []byte, n int) {
	for i:=0; i < n; i++{
		b[i] = 9
	}
}

func p2_after(b []byte, n int) {
	_ = b[n-1]
	for i:=0; i < n; i++{
		b[i] = 9
	}
}
