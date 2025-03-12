package services

import "github.com/Dikontay/hobbyfind/repository/internal/services/cassandra"

type Configs struct {
	Cassandra cassandra.Configs `json:"cassandra"`
}
