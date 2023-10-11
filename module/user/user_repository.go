package user

import (
	"context"
	"golang-clean-architecture/domain/entity"
)

// UserRepository is an interface for user repository contract
type UserRepository interface {
	FindById(ctx context.Context, id string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
	DeleteById(ctx context.Context, id string) error
}
