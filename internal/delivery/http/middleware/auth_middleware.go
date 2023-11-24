package middleware

import (
	"github.com/gofiber/fiber/v2"
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"
)

func NewAuth(userUserCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		userUserCase.Log.Debugf("Authorization : %s", request.Token)

		user, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userUserCase.Log.Debugf("User : %+v", user.ID)
		ctx.Locals("user", user)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *entity.User {
	return ctx.Locals("user").(*entity.User)
}
