package transport

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-architecture/module/contact"
)

type ContactTransportHttp struct {
	ContactUsecase contact.ContactUseCase
}

func NewContactTransportHttp(contactUsecase contact.ContactUseCase) *ContactTransportHttp {
	return &ContactTransportHttp{
		ContactUsecase: contactUsecase,
	}
}

func (receiver *ContactTransportHttp) Create(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello World",
	})
}
