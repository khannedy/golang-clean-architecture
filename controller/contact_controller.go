package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/entity"
	"golang-clean-architecture/model"
	"golang-clean-architecture/usecase"
	"math"
)

type ContactController struct {
	UseCase *usecase.ContactUseCase
	Log     *logrus.Logger
}

func NewContactController(useCase *usecase.ContactUseCase, log *logrus.Logger) *ContactController {
	return &ContactController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *ContactController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	request := new(model.CreateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.UserId = user.ID

	response, err := c.UseCase.Create(user, request)
	if err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

func (c *ContactController) List(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	request := &model.SearchContactRequest{
		UserId: user.ID,
		Name:   ctx.Query("name", ""),
		Email:  ctx.Query("email", ""),
		Phone:  ctx.Query("phone", ""),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(user, request)
	if err != nil {
		c.Log.WithError(err).Error("error searching contact")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.ContactResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *ContactController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)
	contactId := ctx.Params("contactId")

	response, err := c.UseCase.Get(user, contactId)
	if err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

func (c *ContactController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	request := new(model.UpdateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.UserId = user.ID
	request.ID = ctx.Params("contactId")

	response, err := c.UseCase.Update(user, request)
	if err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

func (c *ContactController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)
	contactId := ctx.Params("contactId")

	if err := c.UseCase.Delete(user, contactId); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
