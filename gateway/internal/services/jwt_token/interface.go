package jwt_token

import (
	"github.com/Dikontay/hobbyfind/entities"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(creds entities.User) (string, error)
	CheckToken(token string) (*jwt.Token, error)
}
