package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContactRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}
