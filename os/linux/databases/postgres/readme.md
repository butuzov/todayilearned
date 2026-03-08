# Postgre SQL

## Client Applications

Available tooling:

- https://www.postgresql.org/docs/current/reference-client.html
- [psql](https://www.postgresql.org/docs/current/app-psql.html)
- [pg_dump](https://www.postgresql.org/docs/current/app-pgdump.html)

### `psql`

```shell
psql -U db_user -h localhost db_name
\d
...
> 14 rows affected.
```

Some other examples of [psql](https://www.postgresql.org/docs/current/app-psql.html)

- `\t` - tuples only
- `\a` - no extra breaking elements (added by psql)
- `\q query.sql` - rediect to sql.
- `\o` stop redirection.
- `\i query.sql` - execute generated script
- `\du` - users

## Operations

### Export/Import

```shell

# Export (Interactivly) data
pg_dump -p 5433 -U db_user -h localhost db_name > dump.sql

# Import (Non-Interactivly)
export PGPASSWORD=db_passs
psql -p 5433 -U db_user -h localhost db_name < dump.sql

# Using Container (as tooling home)
docker exec -it postgres-postgresql-1 psql -U db_user -h localhost db_name
```

## Docker

With default image you can do some tricks...

### Importing data on a start.

```yaml
# docker-compose.override.yaml
services:
  db:
    volumes:
      - postgres:/var/lib/postgresql/data/
      - ./northwind.sql:/docker-entrypoint-initdb.d/northwind.sql
```
