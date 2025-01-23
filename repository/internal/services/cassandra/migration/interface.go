package migration

import "github.com/gocql/gocql"

type Service interface {
	Start(session *gocql.Session, keyspace string, dbName string) error
}
