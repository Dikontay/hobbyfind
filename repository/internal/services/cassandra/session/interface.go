package session

import "github.com/gocql/gocql"

type Service interface {
	Init() error
	GetSession() *gocql.Session
	Stop() error
}
