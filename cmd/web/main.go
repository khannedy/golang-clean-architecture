package main

import (
	"fmt"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/internal"
	"golang-clean-architecture/middleware"
	"golang-clean-architecture/usecase"
)

func main() {
	config, err := internal.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	log := internal.NewLogger(config)
	log.Info("Start application")

	db, err := internal.NewDatabase(config, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	validate := internal.NewValidator(config)

	webPort := config.GetInt("web.port")
	app := internal.NewFiber(config)

	routeConfig := internal.RouteConfig{
		App:               app,
		UserController:    controller.NewUserController(usecase.NewUserUseCase(db, log, validate), log),
		ContactController: controller.NewContactController(usecase.NewContactUseCase(db, log, validate), log),
		AddressController: controller.NewAddressController(db, validate, log),
		AuthMiddleware:    middleware.NewAuth(db, log),
	}
	routeConfig.Setup()

	//start server
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatal(err)
	}
}
