// file-header-check
package testdata

import (
	"sync"
)

func fooBarBarBar(a int, b float32, d sync.WaitGroup) { // MATCH /sync.WaitGroup passed by value, the function will get a copy of the original one/
}
