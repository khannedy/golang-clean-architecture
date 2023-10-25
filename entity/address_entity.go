package entity

import "time"

type Address struct {
	ID         string    `gorm:"column:id;primary_key"`
	ContactId  string    `gorm:"column:contact_id"`
	Street     string    `gorm:"column:street"`
	City       string    `gorm:"column:city"`
	Province   string    `gorm:"column:province"`
	PostalCode string    `gorm:"column:postal_code"`
	Country    string    `gorm:"column:country"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
	Contact    Contact   `gorm:"foreignKey:contact_id;references:id"`
}

func (a *Address) TableName() string {
	return "addresses"
}
