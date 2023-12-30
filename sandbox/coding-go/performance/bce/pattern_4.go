package main

func p4(a,b,c []byte) {
	for i := range a {
		a[i] = b[i] + c[i]
	}
}

func p4_after(a,b,c []byte) {
	_ = b[len(a)-1]
	_ = c[len(a)-1]
	for i := range a {
		a[i] = b[i] + c[i]
	}
}
