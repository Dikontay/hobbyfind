package session

type Configs struct {
	Host     string `json:"host"`
	KeySpace string `json:"keyspace"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
