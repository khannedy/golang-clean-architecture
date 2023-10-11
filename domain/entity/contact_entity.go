package entity

import "time"

type Contact struct {
	ID        string    `gorm:"column:id;primary_key"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	UserId    string    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
