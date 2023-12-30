package testdata

func twoReturnParams() (i int, err error) { // want `named return "i" with type "int" found`
	defer func() {
		i = 0
		err = nil
	}()
	return
}
