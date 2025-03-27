package read

import "github.com/gofiber/fiber/v3"

func Endpoint() *fiber.Route {
	return &fiber.Route{
		Method:   "GET",
		Handlers: Handlers,
		Path:     "/read/:id",
	}
}
