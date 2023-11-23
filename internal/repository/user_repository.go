package repository

import (
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/entity"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}
