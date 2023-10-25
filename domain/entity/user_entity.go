package entity

import "time"

// User is a struct that represents a user entity
type User struct {
	ID        string    `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
	Contacts  []Contact `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
