package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
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

	var total int64
	err = tx.Model(&entity.User{}).Where("id = ?", request.ID).Count(&total).Error
	if err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return err
	}

	if total > 0 {
		c.Log.Warnf("User already exists : %+v", err)
		return fiber.ErrConflict
	}

	user := &entity.User{
		ID:       request.ID,
		Password: string(password),
		Name:     request.Name,
	}

	err = tx.Create(user).Error
	if err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return err
	}

	tx.Commit()

	return ctx.JSON(model.WebResponse[model.UserResponse]{Data: model.UserResponse{}})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return err
	}

	err = c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body  : %+v", err)
		return err
	}

	tx := c.DB.Begin()
	defer tx.Rollback()

	user := new(entity.User)
	err = tx.Take(user, "id = ?", request.ID).Error
	if err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return fiber.ErrUnauthorized
	}

	user.Token = uuid.New().String()
	err = tx.Save(user).Error
	if err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return err
	}

	tx.Commit()

	response := model.UserResponse{
		Token: user.Token,
	}

	return ctx.JSON(model.WebResponse[model.UserResponse]{Data: response})
}

func (c *UserController) Current(ctx *fiber.Ctx) error {
	tx := c.DB.Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	user := ctx.Locals("user").(*entity.User)

	response := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.UserResponse]{Data: response})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	tx := c.DB.Begin()
	defer tx.Rollback()

	user := ctx.Locals("user").(*entity.User)
	user.Token = ""

	err := tx.Save(user).Error
	if err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return err
	}

	tx.Commit()

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	tx := c.DB.Begin()
	defer tx.Rollback()

	user := ctx.Locals("user").(*entity.User)

	request := new(model.UpdateUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return err
	}

	err = c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return err
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
			return err
		}
		user.Password = string(password)
	}

	err = tx.Save(user).Error
	if err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return err
	}

	tx.Commit()

	response := model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(model.WebResponse[model.UserResponse]{Data: response})
}
