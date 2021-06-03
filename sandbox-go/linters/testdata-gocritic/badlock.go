package testdata_gocritic

import "sync"

// Example for the rule: badLock

func _() {
	d := 0
	mu := sync.Mutex{}
	mu.Lock()
	d += 1
	mu.Unlock()
	_ = d
}
