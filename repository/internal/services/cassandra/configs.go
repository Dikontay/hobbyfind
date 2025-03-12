package cassandra

import (
	"github.com/Dikontay/hobbyfind/repository/internal/services/cassandra/migration"
	"github.com/Dikontay/hobbyfind/repository/internal/services/cassandra/session"
)

type Configs struct {
	Session   session.Configs   `json:"session"`
	Migration migration.Configs `json:"migration"`
}
