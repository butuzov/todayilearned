package testdata

// PROJECT URL: https://github.com/jingyugao/rowserrcheck

import (
	"database/sql"
)

var db *sql.DB

func RowsErrNotCheck(db *sql.DB) {
	rows, _ := db.Query("") // want "rows.Err must be checked"

	defer func() {
		_ = rows.Close()
	}()
}
