package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"golang-clean-architecture/module/user"
)

type userUseCase struct {
	Validate       *validator.Validate
	UserRepository user.UserRepository
}

func (u *userUseCase) Create(ctx context.Context) error {
	return nil
}
