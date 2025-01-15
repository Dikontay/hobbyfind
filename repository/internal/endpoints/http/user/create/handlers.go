package create

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

var Handlers = []fiber.Handler{create()}

func create() fiber.Handler {
	return func(ctx fiber.Ctx) error {

		body := ctx.Body()
		fmt.Print(body)
		return ctx.SendString("got request")
	}
}
