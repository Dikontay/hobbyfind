package jwt_token

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/entities"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type service struct {
	configs Configs
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
	}
}

func (s *service) CheckToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.configs.JWTSecretKey, nil
	})
}

func (s *service) GenerateToken(creds entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   creds.ID,
		"email": creds.Email,
		"role":  creds.Role,
		"exp":   time.Now().Add(time.Duration(s.configs.JWTExpireTimeInHours) * time.Hour).Unix(),
	})

	return token.SignedString(s.configs.JWTSecretKey)
}
