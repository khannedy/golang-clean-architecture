package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/entity"
	"golang-clean-architecture/middleware"
	"golang-clean-architecture/model"
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

func (c *AddressController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entity.User)

	request := new(model.CreateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = user.ID
	request.ContactId = ctx.Params("contactId")

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", request.ContactId, user.ID).First(contact).Error; err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	address := &entity.Address{
		ID:         uuid.NewString(),
		ContactId:  contact.ID,
		Street:     request.Street,
		City:       request.City,
		Province:   request.Province,
		PostalCode: request.PostalCode,
		Country:    request.Country,
	}

	if err := tx.Create(address).Error; err != nil {
		c.Log.WithError(err).Error("failed to create address")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	response := model.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.AddressResponse]{
		Data: response,
	})
}

func (c *AddressController) List(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", contactId, user.ID).First(contact).Error; err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	var addresses []entity.Address
	if err := tx.Where("contact_id = ?", contact.ID).Find(&addresses).Error; err != nil {
		c.Log.WithError(err).Error("failed to find addresses")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	responses := make([]model.AddressResponse, len(addresses))
	for i, address := range addresses {
		responses[i] = model.AddressResponse{
			ID:         address.ID,
			Street:     address.Street,
			City:       address.City,
			Province:   address.Province,
			PostalCode: address.PostalCode,
			Country:    address.Country,
			CreatedAt:  address.CreatedAt,
			UpdatedAt:  address.UpdatedAt,
		}
	}

	return ctx.JSON(model.WebResponse[[]model.AddressResponse]{
		Data: responses,
	})
}

func (c *AddressController) Get(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	addressId := ctx.Params("addressId")

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", contactId, user.ID).First(contact).Error; err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	address := new(entity.Address)
	if err := tx.Where("id = ? AND contact_id = ?", addressId, contactId).First(address).Error; err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	response := model.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.AddressResponse]{Data: response})
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

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", request.ContactId, user.ID).First(contact).Error; err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	address := new(entity.Address)
	if err := tx.Where("id = ? AND contact_id = ?", request.ID, contact.ID).First(address).Error; err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return fiber.ErrNotFound
	}

	address.Street = request.Street
	address.City = request.City
	address.Province = request.Province
	address.PostalCode = request.PostalCode
	address.Country = request.Country

	if err := tx.Save(address).Error; err != nil {
		c.Log.WithError(err).Error("failed to update address")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	response := model.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.AddressResponse]{Data: response})
}

func (c *AddressController) Delete(ctx *fiber.Ctx) error {
	var (
		user      = middleware.GetUser(ctx)
		contactId = ctx.Params("contactId")
		addressId = ctx.Params("addressId")
	)

	tx := c.DB.Begin()
	defer tx.Rollback()

	contact := new(entity.Contact)
	if err := tx.Where("id = ? AND user_id = ?", contactId, user.ID).First(contact).Error; err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	address := new(entity.Address)
	if err := tx.Where("id = ? AND contact_id = ?", addressId, contactId).First(address).Error; err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return fiber.ErrNotFound
	}

	if err := tx.Delete(address).Error; err != nil {
		c.Log.WithError(err).Error("failed to delete address")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
