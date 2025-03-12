package auth

import "github.com/gofiber/fiber/v3"

func Endpoint() *fiber.Route {
	return &fiber.Route{
		Method:   "POST",
		Handlers: Handlers,
		Path:     "/signup",
	}
}
