package testdata

// OK_with_config
// PROJECT URL:  https://github.com/golangci/golangci-lint/blob/master/pkg/golinters/nolintlint/README.md

import (
	"database/sql"
)

//nolint:wsl,deadcode,unused,gochecknoglobals,varcheck
var db2 *sql.DB

//nolint:noctx,deadcode,unused
func _() {}
