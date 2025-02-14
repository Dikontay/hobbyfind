package create

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"project/internal/domain/models"
	"project/internal/services"
)

var Handlers = []fiber.Handler{parseParams(), create()}

func parseParams() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		body := ctx.Body()
		params := Params{}
		err := json.Unmarshal(body, &params)
		if err != nil {
			return fmt.Errorf("failed to parse params: %v", err)
		}
		ctx.Locals("params", params)
		return ctx.Next()
	}
}

func create() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		params := ctx.Locals("params").(Params)

		user := models.User{
			AuthProperties:     params.AuthProperties,
			UserAdditionalInfo: params.UserAdditionalInfo,
		}

		createdUser, err := services.Repository().CreateUser(user)
		if err != nil {
			return fmt.Errorf("failed to create user: %v", err)
		}

		return ctx.JSON(createdUser)
	}
}
