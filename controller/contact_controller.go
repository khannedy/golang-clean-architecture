package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/entity"
	"golang-clean-architecture/model"
	"gorm.io/gorm"
	"math"
	"strconv"
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
	user := ctx.Locals("user").(*entity.User)

	request := new(model.CreateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.UserId = user.ID
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := &entity.Contact{
		ID:        uuid.New().String(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		UserId:    request.UserId,
	}

	if err := tx.Create(contact).Error; err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return fiber.ErrInternalServerError
	}

	response := model.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.ContactResponse]{Data: response})
}

func (c *ContactController) List(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	tx := c.DB.Begin()
	defer tx.Rollback()

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		c.Log.WithError(err).Error("error parsing page")
		page = 1
	}

	size, err := strconv.Atoi(ctx.Query("size", "10"))
	if err != nil {
		c.Log.WithError(err).Error("error parsing size")
		size = 10
	}

	var contacts []entity.Contact
	if err := tx.Scopes(FilterContact(ctx, user)).Offset((page - 1) * size).Limit(size).Find(&contacts).Error; err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return fiber.ErrInternalServerError
	}

	var total int64 = 0
	if err := tx.Model(&entity.Contact{}).Scopes(FilterContact(ctx, user)).Count(&total).Error; err != nil {
		c.Log.WithError(err).Error("error getting total contacts")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return fiber.ErrInternalServerError
	}

	responses := make([]model.ContactResponse, len(contacts))
	for i, contact := range contacts {
		responses[i] = model.ContactResponse{
			ID:        contact.ID,
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Email:     contact.Email,
			Phone:     contact.Phone,
			CreatedAt: contact.CreatedAt,
			UpdatedAt: contact.UpdatedAt,
		}
	}

	paging := model.PageMetadata{
		Page:      page,
		Size:      size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(size))),
	}

	return ctx.JSON(model.WebResponse[[]model.ContactResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func FilterContact(ctx *fiber.Ctx, user *entity.User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", user.ID)

		name := ctx.Query("name", "")
		if name != "" {
			name = "%" + name + "%"
			tx = tx.Where("first_name LIKE ? OR last_name LIKE ?", name, name)
		}

		phone := ctx.Query("phone", "")
		if phone != "" {
			phone = "%" + phone + "%"
			tx = tx.Where("phone LIKE ?", phone)
		}

		email := ctx.Query("email", "")
		if email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}

		return tx
	}
}

func (c *ContactController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", ctx.Params("contactId"), user.ID).Take(contact).Error; err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return fiber.ErrInternalServerError
	}

	response := model.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.ContactResponse]{Data: response})
}

func (c *ContactController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	request := new(model.UpdateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.UserId = user.ID
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", ctx.Params("contactId"), user.ID).Take(contact).Error; err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return fiber.ErrNotFound
	}

	contact.FirstName = request.FirstName
	contact.LastName = request.LastName
	contact.Email = request.Email
	contact.Phone = request.Phone

	if err := tx.Save(contact).Error; err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return fiber.ErrInternalServerError
	}

	response := model.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.ContactResponse]{Data: response})
}

func (c *ContactController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", ctx.Params("contactId"), user.ID).Take(contact).Error; err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return fiber.ErrNotFound
	}

	if err := tx.Delete(contact).Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
