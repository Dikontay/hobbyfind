package models

type JWTToken struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	Token    string `json:"access_token"`
	ExpireAt int64  `json:"expire_at"`
	Audience string `json:"audience"`
}
