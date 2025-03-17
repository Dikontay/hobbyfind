package jwt_token

type Configs struct {
	JWTSecretKey         string `json:"secret_key"`
	JWTExpireTimeInHours int64  `json:"expire_time_in_hours"`
}
