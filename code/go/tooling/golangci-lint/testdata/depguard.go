package testdata

import "github.com/sirupsen/logrus"

// OK_with_config

// PROJECT URL:  https://github.com/OpenPeeDeeP/depguard
func _() {
	logrus.Debug("I shouldnt be allowed!")
}
