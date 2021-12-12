package main

func p5(b []byte, h []int32) {
	for _, t := range b {
		h[t]++
	}
}

// bytes as they limited, so this is utilize it.
func p5_after(b []byte, h []int32) {

	h = h[:256]

	for _, t := range b {
		h[t]++
	}
}
