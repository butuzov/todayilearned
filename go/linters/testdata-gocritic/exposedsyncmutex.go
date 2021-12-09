package testdata_gocritic

import "sync"

type FooLockable struct {
	sync.Mutex
}
