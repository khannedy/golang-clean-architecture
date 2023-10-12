package user

import "context"

type UserUseCase interface {
	Create(ctx context.Context) error
}
