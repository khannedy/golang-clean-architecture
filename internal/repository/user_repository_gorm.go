package repository

import "gorm.io/gorm"

type userRepositoryGorm struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryGorm{
		DB: db,
	}
}
