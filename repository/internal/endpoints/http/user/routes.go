package user

import (
	"github.com/gofiber/fiber/v3"
	"project/internal/endpoints/http/user/create"
	"project/internal/endpoints/http/user/read"
)

func GetRoutes() []*fiber.Route {

	return []*fiber.Route{
		create.Endpoint(),
		read.Endpoint(),
	}

}
