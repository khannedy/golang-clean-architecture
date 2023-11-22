package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/entity"
	"golang-clean-architecture/middleware"
	"golang-clean-architecture/model"
	"golang-clean-architecture/usecase"
)

type AddressController struct {
	UseCase *usecase.AddressUseCase
	Log     *logrus.Logger
}

func NewAddressController(useCase *usecase.AddressUseCase, log *logrus.Logger) *AddressController {
	return &AddressController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *AddressController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	request := new(model.CreateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = user.ID
	request.ContactId = ctx.Params("contactId")

	response, err := c.UseCase.Create(user, request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

func (c *AddressController) List(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	responses, err := c.UseCase.List(user, contactId)
	if err != nil {
		c.Log.WithError(err).Error("failed to list addresses")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.AddressResponse]{Data: responses})
}

func (c *AddressController) Get(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	addressId := ctx.Params("addressId")

	response, err := c.UseCase.Get(user, contactId, addressId)
	if err != nil {
		c.Log.WithError(err).Error("failed to get address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

func (c *AddressController) Update(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)

	request := new(model.UpdateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = user.ID
	request.ContactId = ctx.Params("contactId")
	request.ID = ctx.Params("addressId")

	response, err := c.UseCase.Update(user, request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update address")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AddressResponse]{Data: response})
}

func (c *AddressController) Delete(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	addressId := ctx.Params("addressId")

	if err := c.UseCase.Delete(user, contactId, addressId); err != nil {
		c.Log.WithError(err).Error("failed to delete address")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
