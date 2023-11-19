package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang-clean-architecture/config"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/internal"
)

func main() {
	viperConfig, err := config.New()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w \n", err))
	}

	log := internal.NewLogger(viperConfig)
	log.Info("Start application")

	db, err := internal.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	validator := internal.NewValidator(viperConfig)

	webPort := viperConfig.GetInt("web.port")
	app := NewFiber(viperConfig)

	//register controller
	userController := controller.NewUserController(db, validator, log)
	userController.Routes(app)
	contactController := controller.NewContactController(db, validator, log)
	contactController.Routes(app)
	addressController := controller.NewAddressController(db, validator, log)
	addressController.Routes(app)

	//start server
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		panic(err)
	}
}

// NewFiber is a function to initialize fiber internal
func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName: config.Get("app.name").(string),
	})

	return app
}
