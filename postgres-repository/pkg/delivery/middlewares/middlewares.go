package middlewares

import (
	"github.com/gofiber/fiber/v3"
)

func IsContentTypeJson() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		// Check Content-Type
		if ctx.Get("Content-Type") != "application/json" {
			return fiber.NewError(fiber.StatusUnsupportedMediaType, "Content-Type must be application/json")
		}
		return ctx.Next()
	}
}
