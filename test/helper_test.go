package test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang-clean-architecture/internal/entity"
	"strconv"
	"testing"
)

func ClearAll() {
	ClearAddresses()
	ClearContact()
	ClearUsers()
}

func ClearUsers() {
	err := db.Where("id is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func ClearContact() {
	err := db.Where("id is not null").Delete(&entity.Contact{}).Error
	if err != nil {
		log.Fatalf("Failed clear contact data : %+v", err)
	}
}

func ClearAddresses() {
	err := db.Where("id is not null").Delete(&entity.Address{}).Error
	if err != nil {
		log.Fatalf("Failed clear address data : %+v", err)
	}
}

func CreateContacts(user *entity.User, total int) {
	for i := 0; i < total; i++ {
		contact := &entity.Contact{
			ID:        uuid.NewString(),
			FirstName: "Contact",
			LastName:  strconv.Itoa(i),
			Email:     "contact" + strconv.Itoa(i) + "@example.com",
			Phone:     "08000000" + strconv.Itoa(i),
			UserId:    user.ID,
		}
		err := db.Create(contact).Error
		if err != nil {
			log.Fatalf("Failed create contact data : %+v", err)
		}
	}
}

func CreateAddresses(t *testing.T, contact *entity.Contact, total int) {
	for i := 0; i < total; i++ {
		address := &entity.Address{
			ID:         uuid.NewString(),
			ContactId:  contact.ID,
			Street:     "Jalan Belum Jadi",
			City:       "Jakarta",
			Province:   "DKI Jakarta",
			PostalCode: "2131323",
			Country:    "Indonesia",
		}
		err := db.Create(address).Error
		assert.Nil(t, err)
	}
}

func GetFirstUser(t *testing.T) *entity.User {
	user := new(entity.User)
	err := db.First(user).Error
	assert.Nil(t, err)
	return user
}

func GetFirstContact(t *testing.T, user *entity.User) *entity.Contact {
	contact := new(entity.Contact)
	err := db.Where("user_id = ?", user.ID).First(contact).Error
	assert.Nil(t, err)
	return contact
}

func GetFirstAddress(t *testing.T, contact *entity.Contact) *entity.Address {
	address := new(entity.Address)
	err := db.Where("contact_id = ?", contact.ID).First(address).Error
	assert.Nil(t, err)
	return address
}
