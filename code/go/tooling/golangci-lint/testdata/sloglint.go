package testdata

import "log/slog"

func _() {
	slog.Info("msg", "foo", 1, slog.Int("bar", 2)) // want `key-value pairs and attributes should not be mixed`
}
