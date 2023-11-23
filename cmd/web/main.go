package main

import (
	"fmt"
	"golang-clean-architecture/internal/config"
	"golang-clean-architecture/internal/delivery/http"
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/delivery/http/route"
	"golang-clean-architecture/internal/usecase"
)

func main() {
	_config, err := config.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	log := config.NewLogger(_config)
	log.Info("Start application")

	db, err := config.NewDatabase(_config, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	validate := config.NewValidator(_config)

	webPort := _config.GetInt("web.port")
	app := config.NewFiber(_config)

	routeConfig := route.RouteConfig{
		App:               app,
		UserController:    http.NewUserController(usecase.NewUserUseCase(db, log, validate), log),
		ContactController: http.NewContactController(usecase.NewContactUseCase(db, log, validate), log),
		AddressController: http.NewAddressController(usecase.NewAddressUseCase(db, log, validate), log),
		AuthMiddleware:    middleware.NewAuth(db, log),
	}
	routeConfig.Setup()

	//start server
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatal(err)
	}
}
