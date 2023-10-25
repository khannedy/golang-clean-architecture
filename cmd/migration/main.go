package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"golang-clean-architecture/config"
	"golang-clean-architecture/internal"
	"gorm.io/gorm"
)

//go:embed migrations
var fs embed.FS

func main() {
	viper, err := config.New()
	if err != nil {
		panic(fmt.Errorf("Fatal error viper file: %w \n", err))
	}

	log := internal.NewLogger(viper)
	log.Info("Start migration")

	db, err := internal.NewDatabase(viper, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	err = RunMigration(db)
	if err != nil {
		panic(fmt.Errorf("Error run migration: %w \n", err))
	}

	log.Info("Finish migration")
}

func RunMigration(db *gorm.DB) error {
	dbSql, err := db.DB()
	if err != nil {
		return err
	}

	location, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(dbSql, &mysql.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("iofs", location, "mysql", driver)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		} else {
			return err
		}
	}

	return nil
}
