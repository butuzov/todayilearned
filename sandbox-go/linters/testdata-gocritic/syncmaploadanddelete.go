package testdata_gocritic

import "golang.org/x/sync/syncmap"

func _(m syncmap.Map, k string) {
	v, ok := m.Load(k)
	if ok {
		m.Delete(k)
		_ = v
	}
}
