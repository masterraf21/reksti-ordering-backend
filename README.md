# Development Setup

## Tools

- Go version 1.3
- Go Modules for dependency management.
- migrate for database migration. https://github.com/golang-migrate/migrate

## Migration

Make sure to use the latest database schema by running:

```shell
migrate \
  -source file://migrations \
  -database "mysql://[username]:[password]@tcp([host]:[port])/[database]" up
```

To modify database schema, make new migration files by using command below:

```shell
migrate create -ext sql -dir migrations [migration_name]
```

There is already script to run migrations in Makefile using test or development configuration

```shell
make migrate-test-up 
```

```shell
make migrate-dev-up 
```

You can change the database configs for migrations

# Run Program

## Without Docker

This program needs environment variables. See example below.

```shell
PORT=8800

READER_HOST=localhost
READER_PORT=3306
READER_USER=root
READER_PASSWORD=password

WRITER_HOST=localhost
WRITER_PORT=3306
WRITER_USER=root
WRITER_PASSWORD=password

MYSQL_DATABASE_NAME=db_name

TIMEOUT_ON_SECONDS=120
OPERATION_ON_EACH_CONTEXT=500
```

You can store environment variables to a file, such as ".env". But, if you want to differentiate environment variables used by docker, you can give another name, such as ".env-no-docker". Run the program with commands below.

```shell
./cmds/env .env go run main.go
```

