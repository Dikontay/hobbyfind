package models

type User struct {
	StandardProperties
	AuthProperties
	UserAdditionalInfo
}

type AuthProperties struct {
	Username string `json:"name" db:"name,unique,notnull"`
	Password string `json:"password" db:"password,notnull"`
}

type UserAdditionalInfo struct {
	Email    string `json:"email" db:"email,unique,notnull"`
	Fullname string `json:"fullname" db:"fullname"`
	Phone    string `json:"phone" db:"phone"`
}
