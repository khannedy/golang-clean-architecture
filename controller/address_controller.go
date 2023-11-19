package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressController struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Log      *logrus.Logger
}

func NewAddressController(db *gorm.DB, validate *validator.Validate, log *logrus.Logger) *AddressController {
	return &AddressController{
		DB:       db,
		Validate: validate,
		Log:      log,
	}
}

func (c *AddressController) Routes(app *fiber.App) {
	app.Get("/api/contacts/:contactId/addresses", c.List)
	app.Post("/api/contacts/:contactId/addresses", c.Create)
	app.Put("/api/contacts/:contactId/addresses/:addressId", c.Update)
	app.Get("/api/contacts/:contactId/addresses/:addressId", c.Get)
	app.Delete("/api/contacts/:contactId/addresses/addressId", c.Delete)
}

func (c *AddressController) Create(ctx *fiber.Ctx) error {
	userContext := ctx.UserContext()
	tx := c.DB.WithContext(userContext).Begin()
	defer tx.Rollback()
	return nil
}

func (c *AddressController) List(ctx *fiber.Ctx) error {
	return nil
}

func (c *AddressController) Get(ctx *fiber.Ctx) error {
	return nil
}

func (c *AddressController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *AddressController) Delete(ctx *fiber.Ctx) error {
	return nil
}
