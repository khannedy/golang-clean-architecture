package test

import (
	"golang-clean-architecture/entity"
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
