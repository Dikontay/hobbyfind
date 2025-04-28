package signup

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/entities"
	"github.com/Dikontay/hobbyfind/gateway/internal/services"
	"github.com/gofiber/fiber/v3"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Handler() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		requestParams := ctx.Context().Value("request_params").(Params)
		hashedPw, err := bcrypt.GenerateFromPassword([]byte(requestParams.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password")
		}

		newUser := entities.User{
			Username: requestParams.Username,
			Role:     requestParams.Role,
			Phone:    requestParams.Phone,
			Email:    requestParams.Email,
			Password: string(hashedPw),
			Fullname: requestParams.FullName,
		}

		createdUser, err := services.Requester().CreateUser(newUser)
		if err != nil {
			log.Warnf("[CLIENTS] repository reponded with error %e", err)
			return ctx.Status(fiber.StatusInternalServerError).SendString("failed to create user")
		}

		token, err := services.JWTToken().GenerateToken(*createdUser)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}

		return ctx.JSON(fiber.Map{
			"user":  createdUser,
			"token": token,
		})

	}

}
