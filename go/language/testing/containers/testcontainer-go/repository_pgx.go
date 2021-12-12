package testcontainer

import (
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type PgXSQLRepository struct {
	db *sqlx.DB
}

func NewPgXSQLRepository(db *sqlx.DB) (PgXSQLRepository, error) {
	return PgXSQLRepository{db}, nil
}

func (repo PgXSQLRepository) Sum(a, b int) int {

	var (
		result struct {
			N int `db:"n"`
		}

		query = "select sum(x::int) as n from (values ($1::NUMERIC), ($2::NUMERIC)) as dt(x)"
	)

	row := repo.db.QueryRowx(query, a, b)
	if err := row.StructScan(&result); err != nil {
		panic(err)
	}

	return result.N
}
