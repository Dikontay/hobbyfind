package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func ParseParams[T any]() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var params T

		// Bind request body to params
		if err := ctx.Bind().Body(params); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{"error": "Failed to parse request body"},
			)
		}

		// If params have a Validate() method, call it
		if validator, ok := any(params).(interface{ Validate() error }); ok {
			if err := validator.Validate(); err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(
					fiber.Map{"error": fmt.Sprintf("Invalid request params: %v", err)},
				)
			}
		}

		// Store parsed params in Locals
		ctx.Locals("request_params", params)
		return ctx.Next()
	}
}
