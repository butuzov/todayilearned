package testdata_gocritic

import (
	"sync"
	"time"
)

// Example for the rule: wrapperFunc

func _(n int) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Add(-1)

		time.Sleep(time.Second)
	}()

	wg.Wait()

}
