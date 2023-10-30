# Golang Clean Architecture Template

## Description

This is golang clean architecture template.

## Architecture

## Tech Stack

- GoFiber (HTTP Framework)
- GORM (ORM)
- MySQL (Database)
- Viper (Configuration)
- Golang Migrate (Database Migration)
- Go Playground Validator (Validation)
- Logrus (Logger)

## Configuration

All configuration is in `config.json` file.

## How to run

### Run application

```bash
go run cmd/web/main.go
```

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://root:@tcp(localhost:3306)/golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

```