package list

import (
	"encoding/json"
	"fmt"
	"github.com/Dikontay/hobbyfind/repository/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

var Handlers = []fiber.Handler{parseParams(), list()}

func parseParams() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		body := ctx.Body()
		params := Params{}
		err := json.Unmarshal(body, &params)
		if err != nil {
			return fmt.Errorf("failed to parse params: %v", err)
		}
		log.Infof("get request to list")
		ctx.Locals("params", params)
		return ctx.Next()
	}
}

func list() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		params := ctx.Locals("params").(Params)

		users, err := services.Repository().ListUsers(params.User)
		if err != nil {
			return fmt.Errorf("failed to list users: %v", err)
		}

		return ctx.JSON(users)
	}
}
