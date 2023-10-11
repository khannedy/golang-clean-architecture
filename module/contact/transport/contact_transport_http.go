package transport

import (
	"belajar-golang-fiber/module/contact"
	"github.com/gofiber/fiber/v2"
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
