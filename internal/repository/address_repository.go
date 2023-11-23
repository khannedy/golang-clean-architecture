package repository

import (
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/entity"
)

type AddressRepository struct {
	Repository[entity.Address]
	Log *logrus.Logger
}
