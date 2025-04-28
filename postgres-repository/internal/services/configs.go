package services

import "github.com/Dikontay/hobbyfind/repository/internal/services/postgres"

type Configs struct {
	Postgres postgres.Configs `json:"postgres"`
}
