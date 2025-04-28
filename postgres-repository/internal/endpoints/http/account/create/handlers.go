package create

import (
	"encoding/json"
	"fmt"
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres"
	"github.com/Dikontay/hobbyfind/repository/pkg/delivery/middlewares"
	"github.com/gofiber/fiber/v3"
)

var Handlers = []fiber.Handler{middlewares.IsContentTypeJson(), create()}

func create() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		requestParams := Params{}
		err := json.Unmarshal(ctx.Body(), &requestParams)
		if err != nil {
			return fmt.Errorf("failed to parse params: %v", err)
		}

		err = requestParams.Validate()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("invalid params %v", err))
		}

		err = postgres.Create(&requestParams.Account)
		if err != nil {
			return fmt.Errorf("failed to create Account: %v", err)
		}

		return ctx.JSON(requestParams.Account)
	}
}
