package main

import (
	"fmt"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/internal"
	"golang-clean-architecture/middleware"
)

func main() {
	viperConfig, err := internal.New()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w \n", err))
	}

	log := internal.NewLogger(viperConfig)
	log.Info("Start application")

	db, err := internal.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	validate := internal.NewValidator(viperConfig)

	webPort := viperConfig.GetInt("web.port")
	app := internal.NewFiber(viperConfig)

	routeConfig := internal.RouteConfig{
		App:               app,
		UserController:    controller.NewUserController(db, validate, log),
		ContactController: controller.NewContactController(db, validate, log),
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
