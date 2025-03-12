package cassandra

import (
	"github.com/Dikontay/hobbyfind/gateway/internal/services/cassandra/migration"
	"github.com/Dikontay/hobbyfind/gateway/internal/services/cassandra/session"
)

type Configs struct {
	Session   session.Configs   `json:"session"`
	Migration migration.Configs `json:"migration"`
}
