package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-architecture/config"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/exception"
	"golang-clean-architecture/repository"
	"golang-clean-architecture/service"
)

func main() {
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	productRepository := repository.NewProductRepository(database)
	productService := service.NewProductService(&productRepository)
	productController := controller.NewProductController(&productService)

	app := fiber.New()
	productController.Route(app)

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
