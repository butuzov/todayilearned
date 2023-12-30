package testdata_gocritic

import "sync"

// Example for the rule: badLock

func _() {
	mu := sync.Mutex{}
	mu.Lock()
	/*! defer is missing, mutex is unlocked immediately */
	mu.Unlock()
}
