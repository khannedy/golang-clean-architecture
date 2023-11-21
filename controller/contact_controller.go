package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContactController struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Log      *logrus.Logger
}

func NewContactController(db *gorm.DB, validate *validator.Validate, log *logrus.Logger) *ContactController {
	return &ContactController{
		DB:       db,
		Validate: validate,
		Log:      log,
	}
}

func (c *ContactController) Create(ctx *fiber.Ctx) error {
	return nil
}

func (c *ContactController) List(ctx *fiber.Ctx) error {
	return nil
}

func (c *ContactController) Get(ctx *fiber.Ctx) error {
	return nil
}

func (c *ContactController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *ContactController) Delete(ctx *fiber.Ctx) error {
	return nil
}