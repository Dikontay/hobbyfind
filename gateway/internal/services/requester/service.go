package requester

import (
	"github.com/Dikontay/hobbyfind/entities"
	"github.com/gofiber/fiber/v2"
)

type service struct {
	configs Configs
	httpClient
}

func NewService(configs Configs) Service {
	return &service{
		configs:    configs,
		httpClient: fiber.AcquireClient(),
	}
}

func (s service) CreateUser(user entities.User) (*entities.User, *int64, error) {
	s.httpClient.Post(s.configs.UserRepositoryUrl).JSON(user).sen

}
