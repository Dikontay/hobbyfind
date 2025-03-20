package services

import (
	"github.com/Dikontay/hobbyfind/gateway/internal/services/jwt_token"
	"github.com/Dikontay/hobbyfind/gateway/internal/services/requester"
)

func Init(configs Configs) error {

	rq = requester.NewService(configs.Requester)
	jwt = jwt_token.NewService(configs.Jwt)
	return nil
}
