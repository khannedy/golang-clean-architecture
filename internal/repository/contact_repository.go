package repository

import (
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/entity"
)

type ContactRepository struct {
	Repository[entity.Contact]
	Log *logrus.Logger
}
