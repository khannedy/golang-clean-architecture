package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang-clean-architecture/cmd/web/route"
	"golang-clean-architecture/config"
	"golang-clean-architecture/internal"
)

func main() {
	viperConfig, err := config.New()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w \n", err))
	}

	log := internal.NewLogger(viperConfig)
	log.Info("Start application")

	_, err = internal.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	webPort := viperConfig.GetInt("web.port")
	app := NewFiber(viperConfig)

	//register routes
	err = route.User(app)
	if err != nil {
		panic(fmt.Errorf("Fatal error route user: %w \n", err))
	}
	err = route.Contact(app)
	if err != nil {
		panic(fmt.Errorf("Fatal error route contact: %w \n", err))
	}
	err = route.Address(app)
	if err != nil {
		panic(fmt.Errorf("Fatal error route address: %w \n", err))
	}

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
