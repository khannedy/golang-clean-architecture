package entity

import "time"

type Address struct {
	ID         string
	ContactId  string
	Street     string
	City       string
	Province   string
	PostalCode string
	Country    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (a *Address) TableName() string {
	return "addresses"
}
