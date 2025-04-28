package read

import (
	"encoding/json"
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres"

	"github.com/gofiber/fiber/v3"
)

var Handlers = []fiber.Handler{read()}

func read() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		requestParams := Params{}

		err := json.Unmarshal(ctx.Body(), &requestParams)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "failed to parse params")
		}
		err = requestParams.Validate()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid params")
		}

		user := &models.User{}
		user.SetId(requestParams.Id)

		err = postgres.Find(user)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}

		return ctx.JSON(user)

	}
}
