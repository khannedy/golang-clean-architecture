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
	"golang-clean-architecture/route"
	"gorm.io/gorm"
)

var app *fiber.App

var db *gorm.DB

var viperConfig *viper.Viper

var log *logrus.Logger

var validate *validator.Validate

func init() {
	var err error

	viperConfig, err = internal.New()
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

	routeConfig := route.RouteConfig{
		App:               app,
		UserController:    controller.NewUserController(db, validate, log),
		ContactController: controller.NewContactController(db, validate, log),
		AddressController: controller.NewAddressController(db, validate, log),
		AuthMiddleware:    middleware.NewAuth(db, log),
	}
	routeConfig.Setup()
}
