package test

import (
	"github.com/google/uuid"
	"golang-clean-architecture/entity"
	"strconv"
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
