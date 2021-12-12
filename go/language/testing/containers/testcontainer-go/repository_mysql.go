package testcontainer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) (MySQLRepository, error) {
	return MySQLRepository{db}, nil
}

func (repo MySQLRepository) Sum(a, b int) int {
	row := repo.db.QueryRow("SELECT ? + ?", a, b)

	var sum int
	if err := row.Scan(&sum); err != nil {
		panic(err)
	}

	return sum
}
