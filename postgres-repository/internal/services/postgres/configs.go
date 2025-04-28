package postgres

import (
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres/connection"
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres/migration"
)

type Configs struct {
	Connection connection.Configs `json:"connection"`
	Migration  migration.Configs  `json:"migration"`
}
