package testdata

// Project URL: https://github.com/daixiang0/gci
import (
	"fmt"

	//nolint:depguard
	_ "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

func _() error {
	return errors.Errorf("%w", fmt.Errorf("ok %s", "1"))
}
