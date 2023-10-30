package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/entity"
	"golang-clean-architecture/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB       *gorm.DB
	Validate *validator.Validate
	Log      *logrus.Logger
}

func NewUserController(db *gorm.DB, validate *validator.Validate, logger *logrus.Logger) *UserController {
	return &UserController{
		DB:       db,
		Validate: validate,
		Log:      logger,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var request model.RegisterUserRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return err
	}

	err = c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return err
	}

	tx := c.DB.Begin()
	defer tx.Rollback()

	user := entity.User{
		ID:       request.ID,
		Password: string(password),
		Name:     request.Name,
	}

	err = tx.Create(&user).Error
	if err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return err
	}

	tx.Commit()

	return ctx.JSON(model.WebResponse[model.UserResponse]{Data: model.UserResponse{}})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var request model.LoginUserRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return err
	}

	err = c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body  : %+v", err)
		return err
	}

	tx := c.DB.Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	var user entity.User
	err = tx.Take(&user, "id = ?", request.ID).Error
	if err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return fiber.ErrUnauthorized
	}

	response := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.UserResponse]{Data: response})
}
