package services

import "github.com/Dikontay/hobbyfind/gateway/internal/services/cassandra"

type Configs struct {
	Cassandra cassandra.Configs `json:"cassandra"`
}
