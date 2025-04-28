package account

import (
	"github.com/Dikontay/hobbyfind/repository/internal/endpoints/http/account/create"
	"github.com/Dikontay/hobbyfind/repository/internal/endpoints/http/account/list"
	"github.com/Dikontay/hobbyfind/repository/internal/endpoints/http/account/read"
	"github.com/gofiber/fiber/v3"
)

func GetRoutes() []*fiber.Route {

	return []*fiber.Route{
		create.Endpoint(),
		read.Endpoint(),
		list.Endpoint(),
	}

}
