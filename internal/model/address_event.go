package model

type AddressEvent struct {
	ID         string `json:"id"`
	ContactId  string `json:"contact_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

func (a *AddressEvent) GetId() string {
	return a.ID
}
