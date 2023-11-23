package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}
