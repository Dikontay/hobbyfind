package session

type Configs struct {
	Host     string `json:"host"`
	KeySpace string `json:"key_space"`
	DbName   string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
