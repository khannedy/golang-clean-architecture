package main

import (
	"belajar-golang-fiber/app"
	"belajar-golang-fiber/domain/entity"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"
	"os"
)

//go:embed migrations
var fs embed.FS

func main() {
	config, err := app.NewConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	db, err := app.NewDatabase(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	// Run migration if argument is migrate
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		err := RunMigration(db)
		if err != nil {
			panic(fmt.Errorf("Error run migration: %w \n", err))
		}
		return
	}

	//fiber := app.NewFiber(config)
	//
	//err = fiber.Listen(":3000")
	//if err != nil {
	//	panic(err)
	//}

	transaction := db.Begin()
	err = SaveUser(transaction)
	if err != nil {
		fmt.Printf("Error save user: %s \n", err.Error())
	}

	connection, _ := db.DB()
	connection.Close()
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
		return err
	}

	return nil
}

func SaveUser(db *gorm.DB) error {
	defer db.Rollback()

	result := db.Create(&entity.User{
		ID:       "2",
		Username: "rahasia",
		Password: "rahasia",
		Name:     "Eko Kurniawan Khannedy",
	})

	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Save User 1 Success")

	result = db.Create(&entity.User{
		ID:       "1",
		Username: "khannedy",
		Password: "rahasia",
		Name:     "Eko Kurniawan Khannedy",
	})

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Save User 2 Success")

	result = db.Commit()
	if result.Error != nil {
		return result.Error
	}

	return nil
}
