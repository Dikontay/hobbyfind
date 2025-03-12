package models

type User struct {
	StandardProperties
	Username string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
}
