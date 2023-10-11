package entity

import "time"

// User is a struct that represents a user entity
type User struct {
	ID        string    `gorm:"column:id" gorm:"primary_key"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"column:updated_at" gorm:"autoCreateTime:true" gorm:"autoUpdateTime:true"`
}

func (u *User) TableName() string {
	return "users"
}
