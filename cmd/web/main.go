package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang-clean-architecture/config"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/internal"
	"golang-clean-architecture/middleware"
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
	contactController := controller.NewContactController(db, validator, log)
	addressController := controller.NewAddressController(db, validator, log)

	// guest routes
	app.Post("/api/users", userController.Register)
	app.Post("/api/users/_login", userController.Login)

	// auth routes
	app.Use(middleware.NewAuth(db, log))
	app.Delete("/api/users", userController.Logout)
	app.Patch("/api/users/_current", userController.Update)
	app.Get("/api/users/_current", userController.Current)

	app.Get("/api/contacts", contactController.List)
	app.Post("/api/contacts", contactController.Create)
	app.Put("/api/contacts/:contactId", contactController.Update)
	app.Get("/api/contacts/:contactId", contactController.Get)
	app.Delete("/api/contacts/:contactId", contactController.Delete)

	app.Get("/api/contacts/:contactId/addresses", addressController.List)
	app.Post("/api/contacts/:contactId/addresses", addressController.Create)
	app.Put("/api/contacts/:contactId/addresses/:addressId", addressController.Update)
	app.Get("/api/contacts/:contactId/addresses/:addressId", addressController.Get)
	app.Delete("/api/contacts/:contactId/addresses/addressId", addressController.Delete)

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
