package session

import (
	"fmt"
	"github.com/gocql/gocql"
)

type service struct {
	configs Configs
	session *gocql.Session
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
	}
}

func (s *service) Init() error {
	cluster := gocql.NewCluster(s.configs.Host)
	cluster.Port = s.configs.Port
	cluster.Keyspace = s.configs.KeySpace
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: s.configs.Username,
		Password: s.configs.Password,
	}

	var err error

	s.session, err = cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	return nil

}

// todo COMPLETE IT
func (s *service) GetSession() *gocql.Session {
	if s.session == nil {
		panic("session is not initialized")
	}
	return s.session
}

func (s *service) Stop() error {
	s.session.Close()
	return nil
}
