package list

import (
	"encoding/json"
	"fmt"
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres"
	"github.com/gofiber/fiber/v3"
)

var Handlers = []fiber.Handler{list()}

func list() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		requestParams := Params{}

		err := json.Unmarshal(ctx.Body(), &requestParams)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "failed to parse params")
		}

		err = requestParams.Validate()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("invalid params: %v", err))

		}

		users, err := postgres.List(requestParams.User, requestParams.PageNum, requestParams.PageSize)
		if err != nil {
			return fmt.Errorf("failed to list users: %v", err)
		}

		return ctx.JSON(users)
	}
}
