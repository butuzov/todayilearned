package testdata

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db3 *sqlx.DB

// OK_with_config
// PROJECT URL:  https://github.com/ryanrolds/sqlclosecheck

func missingClose() {
	age := 27
	rows, err := db3.Queryx("SELECT name FROM users WHERE age=?", age)
	if err != nil {
		log.Fatal(err)
	}

	// defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	_ = names

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s are %d years old", strings.Join(names, ", "), age)
}
