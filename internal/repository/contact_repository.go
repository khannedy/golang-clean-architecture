package repository

import (
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
	"gorm.io/gorm"
)

type ContactRepository struct {
	Repository[entity.Contact]
	Log *logrus.Logger
}

func NewContactRepository(log *logrus.Logger) *ContactRepository {
	return &ContactRepository{
		Log: log,
	}
}

func (r *ContactRepository) FindByIdAndUserId(db *gorm.DB, contact *entity.Contact, id string, userId string) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(contact).Error
}

func (r *ContactRepository) Search(db *gorm.DB, request *model.SearchContactRequest) ([]entity.Contact, int64, error) {
	var contacts []entity.Contact
	if err := db.Scopes(r.FilterContact(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&contacts).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Contact{}).Scopes(r.FilterContact(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}

func (r *ContactRepository) FilterContact(request *model.SearchContactRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", request.UserId)

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("first_name LIKE ? OR last_name LIKE ?", name, name)
		}

		if phone := request.Phone; phone != "" {
			phone = "%" + phone + "%"
			tx = tx.Where("phone LIKE ?", phone)
		}

		if email := request.Email; email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}

		return tx
	}
}
