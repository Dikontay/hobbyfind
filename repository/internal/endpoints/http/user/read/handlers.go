package read

import (
	"github.com/Dikontay/hobbyfind/repository/internal/services"
	"github.com/Dikontay/hobbyfind/repository/utils"
	"github.com/gofiber/fiber/v3"
)

var Handlers = []fiber.Handler{read()}

func read() fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")

		if !utils.CheckUuid(id) {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
		}

		user, err := services.Repository().ReadUser(id)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}

		return c.JSON(user)

	}
}
