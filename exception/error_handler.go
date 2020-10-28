package exception

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-architecture/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	validationError, ok := err.(ValidationError)
	if !ok {
		ctx.JSON(model.WebResponse{
			Code:   500,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	} else {
		ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   validationError.Error(),
		})
	}

	return nil
}
