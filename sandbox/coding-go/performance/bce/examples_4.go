// Example from the go101 article.
package main


func fd(is []int, bs []byte) {
	if len(is) >= 256 {
		for _, n := range bs {
			_ = is[n] // line 7: bounds check
		}
	}
}

func fd2(is []int, bs []byte) {
	if len(is) >= 256 {
		is = is[:256] // line 14: a hint
		for _, n := range bs {
			_ = is[n] // line 16: BCEed!
		}
	}
}

func fe_before(isa []int, isb []int) {
	if len(isa) > 0xFFF {
		for _, n := range isb {
			_ = isa[n & 0xFFF] // line 24: bounds check
		}
	}
}

func fe_after(isa []int, isb []int) {
	if len(isa) > 0xFFF {
		isa = isa[:0xFFF+1] // line 31: a hint
		for _, n := range isb {
			_ = isa[n & 0xFFF] // line 33: BCEed!
		}
	}
}
