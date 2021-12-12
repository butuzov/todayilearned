package testcontainer

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type (
	PostgreSettings struct {
		User  string
		Pass  string
		Host  string
		Ports []string
		Name  string
		Extra string

		terminate func()
	}
)

// Step 1: Define Setup and Terminate on db settings.

func (db *PostgreSettings) Setup() (*sqlx.DB, error) {
	ctx := context.Background()

	pg, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:11.4-alpine",
			ExposedPorts: db.Ports,
			Env: map[string]string{
				"POSTGRES_PASSWORD": db.Pass,
				"POSTGRES_USER":     db.User,
				"POSTGRES_DB":       db.Name,
			},
			WaitingFor: wait.ForAll(
				wait.ForLog("database system is ready to accept connections"),
				wait.ForListeningPort(nat.Port(db.Ports[0])),
			),
		},
		Started: true,
	})
	if err != nil {
		return nil, err
	}

	db.terminate = func() {
		if err != pg.Terminate(ctx) {
			log.Fatalf("error terminating mysql container: %s", err)
		}
	}

	var (
		host, _ = pg.Host(ctx)
		port, _ = pg.MappedPort(ctx, nat.Port(db.Ports[0]))

		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
			db.User, db.Pass, host, port.Port(), db.Name, db.Extra)

		// using this (canonical) dsn is also OK!
		// dsn = fmt.Sprintf("user=%[1]s password=%[2]s host=%[3]s port=%[4]s dbname=%[5]s %[6]s",
		// 	db.User, db.Pass, host, port.Port(), db.Name, db.Extra)
	)

	dbx, err := sqlx.Connect("postgres", dsn)
	if err == nil {
		return dbx, nil
	}

	return nil, fmt.Errorf("error connect to db: %+v\n", err)
}

func (db PostgreSettings) Terminate() {
	if db.terminate != nil {
		db.terminate()
	}
}

// Step 2: Crate Initial Test for Postgres using PQ driver.

func Test_PostgressRepository(t *testing.T) {
	pg := PostgreSettings{
		User:  "db_user",
		Pass:  "db_pass",
		Host:  "localhost",
		Ports: []string{"5432/tcp"},
		Name:  "db_name",
		Extra: "sslmode=disable",
	}

	if db, err := pg.Setup(); err == nil {
		t.Cleanup(pg.Terminate)

		if repo, err := NewPostgreSQLRepository(db); err == nil {
			suite.Run(t, &RepositoryTestSuite{Repo: repo})
		} else {
			assert.Fail(t, err.Error())
		}

	} else {
		assert.Fail(t, err.Error())
	}
}
