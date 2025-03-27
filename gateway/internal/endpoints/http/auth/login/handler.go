package login

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/entities"
	"github.com/Dikontay/hobbyfind/gateway/internal/services"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func Handler() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		requestParams := ctx.Context().Value("request_params").(Params)

		userProperties := entities.User{
			Email: requestParams.Email,
		}
		users, err := services.Requester().ListUsers(userProperties)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fmt.Errorf("failed to get list of users %v", err))
		}
		if len(users) == 0 {
			return ctx.Status(fiber.StatusNotFound).JSON(fmt.Errorf("user not found %v", err))
		}

		err = bcrypt.CompareHashAndPassword([]byte(requestParams.Password), []byte(users[0].Password))
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		token, err := services.JWTToken().GenerateToken(users[0])
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login successful",
			"token":   token,
		},
		)
	}

}
