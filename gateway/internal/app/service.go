package app

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/gateway/internal/endpoints/http/user"
	"github.com/Dikontay/hobbyfind/gateway/internal/services"
	"github.com/Dikontay/hobbyfind/gateway/utils"
	"github.com/gofiber/fiber/v3"
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

	servicesConfigs := services.Configs{}

	err := utils.InitConfigs("./configs/configs.json", &servicesConfigs)
	if err != nil {
		return fmt.Errorf("failed to init configs: %v", err)
	}

	err = services.Init(servicesConfigs)
	if err != nil {
		return fmt.Errorf("failed to init services: %v", err)
	}
	return s.app.Listen(s.configs.Port)
}

func (s service) Stop() error {
	//TODO implement me
	panic("implement me")
}
