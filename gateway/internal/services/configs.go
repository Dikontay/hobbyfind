package services

import (
	"github.com/Dikontay/hobbyfind/gateway/internal/services/jwt_token"
	"github.com/Dikontay/hobbyfind/gateway/internal/services/requester"
)

type Configs struct {
	Jwt       jwt_token.Configs `json:"jwt"`
	Requester requester.Configs `json:"requester"`
}
