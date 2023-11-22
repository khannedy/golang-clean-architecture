package test

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang-clean-architecture/controller"
	"golang-clean-architecture/internal"
	"golang-clean-architecture/middleware"
	"golang-clean-architecture/usecase"
	"gorm.io/gorm"
)

var app *fiber.App

var db *gorm.DB

var viperConfig *viper.Viper

var log *logrus.Logger

var validate *validator.Validate

func init() {
	var err error

	viperConfig, err = internal.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error viperConfig file: %w \n", err))
	}

	log = internal.NewLogger(viperConfig)
	validate = internal.NewValidator(viperConfig)
	app = internal.NewFiber(viperConfig)

	db, err = internal.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	routeConfig := internal.RouteConfig{
		App:               app,
		UserController:    controller.NewUserController(usecase.NewUserUseCase(db, log, validate), log),
		ContactController: controller.NewContactController(usecase.NewContactUseCase(db, log, validate), log),
		AddressController: controller.NewAddressController(usecase.NewAddressUseCase(db, log, validate), log),
		AuthMiddleware:    middleware.NewAuth(db, log),
	}
	routeConfig.Setup()
}
