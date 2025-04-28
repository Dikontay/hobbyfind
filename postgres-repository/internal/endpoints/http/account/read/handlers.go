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

		Account := &models.Account{}
		Account.SetId(requestParams.Id)

		err = postgres.Find(Account)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Account not found")
		}

		return ctx.JSON(Account)

	}
}
