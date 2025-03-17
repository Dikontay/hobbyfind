package signup

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

var handlers = []fiber.Handler{
	parseBody(),
	hashPassword(),
}

func parseBody() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		body := new(Params)
		err := ctx.Bind().Body(body)
		if err != nil {
			return fmt.Errorf("failed to bind request body")
		}
		err = body.Validate()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				fmt.Errorf("invalid request params %v", err),
			)
		}
		ctx.Locals("request_params", body)
		return ctx.Next()
	}
}

func hashPassword() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		requestParams := ctx.Context().Value("params")
		hashedPw, err := bcrypt.GenerateFromPassword([]byte(requestParams.), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password")
		}
		body.Password = string(hashedPw)

		return ctx.Next()
	}
}
