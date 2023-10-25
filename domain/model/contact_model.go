package model

import "time"

type ContactResponse struct {
	ID        string            `json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Addresses []AddressResponse `json:"addresses"`
}

type CreateContactRequest struct {
	UserId    string `json:"-" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}
