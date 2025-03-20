package services

import (
	"github.com/Dikontay/hobbyfind/gateway/internal/services/jwt_token"
	"github.com/Dikontay/hobbyfind/gateway/internal/services/requester"
)

//singleton private services

var rq requester.Service
var jwt jwt_token.Service

func Requester() requester.Service {
	return rq
}
func JWTToken() jwt_token.Service {
	return jwt
}
