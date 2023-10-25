package transport

import "github.com/sirupsen/logrus"

func NewAddressController(logger *logrus.Logger) *AddressController {
	return &AddressController{
		log: logger,
	}
}

type AddressController struct {
	log *logrus.Logger
}
