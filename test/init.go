package test

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang-clean-architecture/internal/config"
	"golang-clean-architecture/internal/delivery/http"
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/delivery/http/route"
	"golang-clean-architecture/internal/usecase"
	"gorm.io/gorm"
)

var app *fiber.App

var db *gorm.DB

var viperConfig *viper.Viper

var log *logrus.Logger

var validate *validator.Validate

func init() {
	var err error

	viperConfig, err = config.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w \n", err))
	}

	log = config.NewLogger(viperConfig)
	validate = config.NewValidator(viperConfig)
	app = config.NewFiber(viperConfig)

	db, err = config.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	routeConfig := route.RouteConfig{
		App:               app,
		UserController:    http.NewUserController(usecase.NewUserUseCase(db, log, validate), log),
		ContactController: http.NewContactController(usecase.NewContactUseCase(db, log, validate), log),
		AddressController: http.NewAddressController(usecase.NewAddressUseCase(db, log, validate), log),
		AuthMiddleware:    middleware.NewAuth(db, log),
	}
	routeConfig.Setup()
}
