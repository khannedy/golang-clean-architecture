package entity

import "time"

type Contact struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Contact) TableName() string {
	return "contacts"
}
