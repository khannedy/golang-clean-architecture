package repository

import (
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/entity"
	"gorm.io/gorm"
)

type AddressRepository struct {
	Log *logrus.Logger
}

func (r *AddressRepository) Create(db *gorm.DB, address *entity.Address) error {
	return db.Create(address).Error
}
