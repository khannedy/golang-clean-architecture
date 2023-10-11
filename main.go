package main

import (
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"golang-clean-architecture/app"
	"golang-clean-architecture/domain/entity"
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

	log := app.NewLogger(config)
	log.Info("Start application")

	db, err := app.NewDatabase(config, log)
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
	//err = SaveContact(transaction)
	//if err != nil {
	//	fmt.Printf("Error save user: %s \n", err.Error())
	//}

	users, err := FindUserWithContact(transaction)
	if err != nil {
		fmt.Printf("Error find user: %s \n", err.Error())
	}
	log.Info(users)

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

func FindUserWithContact(db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	err := db.Model(&entity.User{}).Preload("Contacts").Find(&users).Error
	return users, err
}

func SaveContact(db *gorm.DB) error {
	defer db.Rollback()

	err := db.Create(&entity.Contact{
		ID:        "1",
		FirstName: "Eko Kurniawan",
		LastName:  "Khannedy",
		Email:     "eko@gmail.com",
		Phone:     "32424234",
		UserId:    "1",
	}).Error
	if err != nil {
		return err
	}

	err = db.Create(&entity.Contact{
		ID:        "2",
		FirstName: "Eko Kurniawan",
		LastName:  "Khannedy",
		Email:     "eko@gmail.com",
		Phone:     "32424234",
		UserId:    "1",
	}).Error
	if err != nil {
		return err
	}

	err = db.Commit().Error
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
