package testdata

import "time"

// PROJECT URL:
// https://github.com/charithe/durationcheck

func durationcheck() {
	type myStruct struct {
		fieldA int
		fieldB time.Duration
		fieldC *int
	}

	x := 30 * time.Second
	ms := myStruct{fieldA: 10, fieldB: 10 * time.Second}
	SomeDuration := 10 * time.Second
	const timeout = 10

	f := x * time.Second
	_ = f
	_ = time.Second * x
	_ = timeout * time.Millisecond
	_ = (30 * time.Second) * time.Millisecond
	_ = time.Millisecond * (30 * time.Second)
	_ = time.Millisecond * time.Second * 1
	_ = 1 * time.Second * (time.Second)
	_ = ms.fieldB * time.Second
	_ = time.Second * ms.fieldB
	_ = SomeDuration * time.Second
	_ = time.Second * SomeDuration
}
