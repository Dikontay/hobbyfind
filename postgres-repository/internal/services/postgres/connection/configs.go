package connection

type Configs struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
	Schema   string `json:"schema"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SslMode  string `json:"ssl_mode"`
}
