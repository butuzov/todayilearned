package testdata

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// PROJECT URL: https://github.com/gordonklaus/ineffassign

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
)

func _() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err = db.Ping(); err != nil {
		println(err)
	}

	insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
	db.Exec(insertStmt)

	rows, _ := db.Query("select id from tb")
	for rows.Next() {
		// handle rows
	}
}
