package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// NewFiber is a function to initialize fiber app
func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName: config.Get("app.name").(string),
	})

	return app
}
