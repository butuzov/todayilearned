# `testcontainer-go`

* https://github.com/testcontainers/testcontainers-go
* https://golang.testcontainers.org/

## About Examples

* We using Setups and Teardowns (so we using test suite)
* We actually testing minimalistic implementation of repository pattern (see DDD and Ports and Adapters). And testing different adapters with out port (Repository)


## Suite

{{% list "suite_test.go,repository.go" %}}

## MySQL

{{% list "repository_mysql.go,repository_mysql_test.go" %}}

## Postgress (& `pgx`)

{{% list "repository_pgx.go,repository_pgx_test.go,repository_postgres.go,repository_postgres_test.go" %}}

## Extras

{{% list "Makefile,docker-compose.yml" %}}
