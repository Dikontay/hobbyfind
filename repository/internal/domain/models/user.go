package models

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `swaggerignore:"true" bun:"table:users,alias:u"`
	StandardProperties
	AuthProperties
	UserAdditionalInfo
}

type AuthProperties struct {
	Username string `json:"name" bun:"name,unique,notnull"`
	Password string `json:"password" bun:"password,notnull"`
}

type UserAdditionalInfo struct {
	Email    string `json:"email" bun:"email,unique,notnull"`
	Fullname string `json:"fullname" bun:"fullname"`
	Phone    string `json:"phone" bun:"phone"`
}
