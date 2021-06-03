package testdata

import "github.com/sirupsen/logrus"

// PROJECT URL:  https://github.com/OpenPeeDeeP/depguard
func _() {
	logrus.Debug("I shouldnt be allowed!")
}
