package main

func p3(b []byte, n int) {

	_ = b[n+5]

	b[n] = byte(1 >> 40)
	b[n+1] = byte(1 >> 32)
	b[n+2] = byte(1 >> 24)
	b[n+3] = byte(1 >> 16)
	b[n+4] = byte(1 >> 8)
	b[n+5] = byte(1)

}

// using re slicing.

func p3_after(b []byte, n int) {

	b = b[n: n+5]

	b[0] = byte(1 >> 40)
	b[1] = byte(1 >> 32)
	b[2] = byte(1 >> 24)
	b[3] = byte(1 >> 16)
	b[4] = byte(1 >> 8)
	b[5] = byte(1)

}
