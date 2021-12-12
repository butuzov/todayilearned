package testcontainer

import (
	"context"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type (
	MySQLSettings struct {
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

func (db *MySQLSettings) Setup() (*sqlx.DB, error) {
	var (
		ctx = context.Background()
		req = testcontainers.ContainerRequest{
			Image:        "mysql:latest",
			ExposedPorts: db.Ports,
			Env: map[string]string{
				"MYSQL_ROOT_PASSWORD": db.Pass,
				"MYSQL_DATABASE":      db.Name,
				"MYSQL_USER":          db.User,
				"MYSQL_PASSWORD":      db.Pass,
			},
			WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
		}
	)

	mysql, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	db.terminate = func() {
		if err != mysql.Terminate(ctx) {
			log.Fatalf("error terminating mysql container: %s", err)
		}
	}

	var (
		host, _ = mysql.Host(ctx)
		port, _ = mysql.MappedPort(ctx, nat.Port(db.Ports[0]))

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			db.User, db.Pass, host, port.Port(), db.Name, db.Extra)
	)

	dbx, err := sqlx.Connect("mysql", dsn)
	if err == nil {
		return dbx, nil
	}

	return nil, fmt.Errorf("error connect to db: %+v\n", err)
}

func (db MySQLSettings) Terminate() {
	if db.terminate != nil {
		db.terminate()
	}
}

// Step 2: Crate Initial Test for MySQL.

func Test_MySQLRepository(t *testing.T) {
	mysql := MySQLSettings{
		User:  "db_user",
		Pass:  "db_pass",
		Host:  "localhost",
		Ports: []string{"3306/tcp", "33060/tcp"},
		Name:  "db_name",
		Extra: "tls=skip-verify&parseTime=true&multiStatements=true",
	}

	if db, err := mysql.Setup(); err == nil {

		t.Cleanup(mysql.Terminate)

		if repo, err := NewMySQLRepository(db); err == nil {
			suite.Run(t, &RepositoryTestSuite{Repo: repo})
		} else {
			assert.Fail(t, err.Error())
		}

	} else {
		assert.Fail(t, err.Error())
	}
}
