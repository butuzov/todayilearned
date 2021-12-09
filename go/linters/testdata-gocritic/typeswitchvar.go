package testdata_gocritic

// Example for the rule: typeSwitchVar

type (
	point struct {
		x, y int
	}
)

func _(v interface{}) interface{} {

	switch v.(type) {
	case int:
		return v.(int)
	case point:
		return v.(point).x + v.(point).y
	default:
		return 0
	}

}
