// file-header-check
package testdata

type tt int

func (t tt) m() {}

func deferrer() {
	for {
		go func() {
			defer println()
		}()
		defer func() {}() // MATCH /prefer not to defer inside loops/
	}
}
