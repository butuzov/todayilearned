package testcontainer

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQLRepository struct {
	db *sqlx.DB
}

func NewPostgreSQLRepository(db *sqlx.DB) (PostgreSQLRepository, error) {
	return PostgreSQLRepository{db}, nil
}

func (repo PostgreSQLRepository) Sum(a, b int) int {
	row := repo.db.QueryRow("select sum(x::int) from (values ($1), ($2)) as dt(x)", a, b)

	var sum int
	if err := row.Scan(&sum); err != nil {
		panic(err)
	}

	return sum
}
