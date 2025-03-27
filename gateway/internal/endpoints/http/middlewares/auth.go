package middlewares

import (
	"github.com/Dikontay/hobbyfind/entities"
	"github.com/Dikontay/hobbyfind/gateway/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func JWT() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		headers := ctx.GetReqHeaders()
		authHeader, ok := headers["Authorization"]
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}
		if authHeader[0] == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		token, err := services.JWTToken().CheckToken(authHeader[0])
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid jwt token"})
		}
		claims := token.Claims.(jwt.MapClaims)
		userID, ok := claims["sub"].(string)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user role"})
		}

		user := entities.User{
			StandardProperties: entities.StandardProperties{
				ID: userID,
			},
			Role: role,
		}

		ctx.Locals("user", user)

		return ctx.Next()
	}
}
