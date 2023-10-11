package datasource

import (
	"context"
	"golang-clean-architecture/domain/entity"
	"golang-clean-architecture/module/user"
)

type userRepositoryMySQL struct {
}

func NewMySQL() user.UserRepository {
	return &userRepositoryMySQL{}
}

func (u *userRepositoryMySQL) FindById(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryMySQL) FindAll(ctx context.Context) ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryMySQL) Save(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryMySQL) Update(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryMySQL) Delete(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryMySQL) DeleteById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
