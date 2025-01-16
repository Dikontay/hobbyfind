package app

import (
	"github.com/gofiber/fiber/v3"
	"project/internal/endpoints/http/user"
)

type service struct {
	configs Configs
	app     *fiber.App
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
		app:     fiber.New(configs.App),
	}
}

func (s service) SetupRoutes() {

	userRoutes := user.GetRoutes()
	for _, route := range userRoutes {
		s.app.Group(s.configs.BasePath).Group(route.Path, route.Handlers...)
	}
	return
}

func (s service) Start() error {
	s.SetupRoutes()
	return s.app.Listen(s.configs.Port)
}

func (s service) Stop() error {
	//TODO implement me
	panic("implement me")
}
