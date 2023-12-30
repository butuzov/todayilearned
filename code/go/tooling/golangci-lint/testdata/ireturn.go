package testdata

type Doer interface {
	Do()
}

type do struct{}

func (d *do) Do() {}

// PROJECT URL:  https://github.com/butuzov/ireturn
func _() Doer {
	return &do{}
}
