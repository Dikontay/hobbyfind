package user

import (
	"github.com/Dikontay/hobbyfind/repository/internal/endpoints/http/user/create"
	"github.com/Dikontay/hobbyfind/repository/internal/endpoints/http/user/list"
	"github.com/Dikontay/hobbyfind/repository/internal/endpoints/http/user/read"
	"github.com/gofiber/fiber/v3"
)

func GetRoutes() []*fiber.Route {

	return []*fiber.Route{
		create.Endpoint(),
		read.Endpoint(),
		list.Endpoint(),
	}

}
