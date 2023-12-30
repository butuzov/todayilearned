package testdata_gocritic

import (
	"database/sql"
	"fmt"
)

func _(host, port, user, password, dbname string) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err = db.Ping(); err != nil {
		println(err)
	}

	_, _ = db.Query("UPDATE tb SET d=1")
}
