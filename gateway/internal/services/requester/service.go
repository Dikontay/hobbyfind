package requester

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/entities"
	"github.com/go-resty/resty/v2"
)

type service struct {
	configs    Configs
	httpClient *resty.Client
}

func NewService(configs Configs) Service {
	return &service{
		configs:    configs,
		httpClient: new(resty.Client),
	}
}

func (s service) CreateUser(user entities.User) (*entities.User, error) {
	result := new(entities.User)
	resp, err := s.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		SetResult(result). // Автоматически маппит JSON в структуру
		Post(s.configs.UserRepositoryUrl)

	if err != nil {
		return nil, err
	}

	// Если код ответа не 200 или 201 — ошибка
	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		return nil, fmt.Errorf("clients repository responded with %d and error %e", resp.StatusCode(), err)
	}

	return result, nil

}

func (s service) ListUsers(user entities.User) ([]entities.User, error) {
	var result []entities.User
	resp, err := s.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(user).
		SetResult(result). // Автоматически маппит JSON в структуру
		Post(fmt.Sprintf("%s/list", s.configs.UserRepositoryUrl))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("clients repository responded with %d and error %e", resp.StatusCode(), err)
	}
	return result, nil

}
