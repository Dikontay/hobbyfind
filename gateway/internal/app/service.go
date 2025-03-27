package app

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/gateway/internal/endpoints/http/auth/login"
	"github.com/Dikontay/hobbyfind/gateway/internal/endpoints/http/auth/signup"
	"github.com/Dikontay/hobbyfind/gateway/internal/endpoints/http/middlewares"
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
	router := s.app.Group("api")

	router.Post("/auth/login",
		login.Handler(),
		middlewares.ParseParams[login.Params](),
	)
	router.Post("/auth/signup",
		signup.Handler(),
		middlewares.ParseParams[signup.Params](),
	)

	return
}

func (s service) Start() error {

	servicesConfigs := services.Configs{}

	err := utils.InitConfigs("./configs/configs.json", &servicesConfigs)
	if err != nil {
		return fmt.Errorf("failed to init configs: %v", err)
	}

	err = services.Init(servicesConfigs)
	if err != nil {
		return fmt.Errorf("failed to init services: %v", err)
	}
	s.SetupRoutes()
	return s.app.Listen(s.configs.Port)
}

func (s service) Stop() error {
	//TODO implement me
	panic("implement me")
}
