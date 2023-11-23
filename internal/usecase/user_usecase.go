package usecase

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate) *UserUseCase {
	return &UserUseCase{
		DB:       db,
		Log:      logger,
		Validate: validate,
	}
}

func (c *UserUseCase) Create(request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := c.DB.Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	var total int64
	err = tx.Model(&entity.User{}).Where("id = ?", request.ID).Count(&total).Error
	if err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		c.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	user := &entity.User{
		ID:       request.ID,
		Password: string(password),
		Name:     request.Name,
	}

	err = tx.Create(user).Error
	if err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	tx.Commit()

	response := &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return response, nil
}

func (c *UserUseCase) Login(request *model.LoginUserRequest) (*model.UserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body  : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	tx := c.DB.Begin()
	defer tx.Rollback()

	user := new(entity.User)
	if err := tx.Take(user, "id = ?", request.ID).Error; err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	user.Token = uuid.New().String()
	if err := tx.Save(user).Error; err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	tx.Commit()

	response := &model.UserResponse{
		Token: user.Token,
	}

	return response, nil
}

func (c *UserUseCase) Current(user *entity.User) (*model.UserResponse, error) {
	tx := c.DB.Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	response := &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return response, nil
}

func (c *UserUseCase) Logout(user *entity.User) (bool, error) {
	tx := c.DB.Begin()
	defer tx.Rollback()

	user.Token = ""

	if err := tx.Save(user).Error; err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

func (c *UserUseCase) Update(user *entity.User, request *model.UpdateUserRequest) (*model.UserResponse, error) {
	tx := c.DB.Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
			return nil, fiber.ErrInternalServerError
		}
		user.Password = string(password)
	}

	if err := tx.Save(user).Error; err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	response := &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
