package testdata

import "fmt"

// OK_with_config
// PROJECT URL:  https://github.com/ashanbrown/forbidigo
func forbidigo() error {
	return fmt.Errorf("forbidigo %v", "enabled")
}
