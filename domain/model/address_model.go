package model

import "time"

type AddressResponse struct {
	ID         string    `json:"id"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
