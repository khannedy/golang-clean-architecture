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

## Configuration

All configuration is in `config.json` file.

## How to run

### Run database migration

```bash
go run main.go migrate
```

### Run application

```bash
go run main.go
```