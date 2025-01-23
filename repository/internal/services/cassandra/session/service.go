package session

import (
	"fmt"
	"github.com/gocql/gocql"
)

type service struct {
	configs Configs
	*gocql.Session
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
	}
}

func (s service) Init() error {
	cluster := gocql.NewCluster(s.configs.Host)
	cluster.Keyspace = s.configs.KeySpace
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: s.configs.Username,
		Password: s.configs.Password,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}
	s.Session = session
	return nil

}
func (s service) GetSession() *gocql.Session {
	return s.Session
}

func (s service) Stop() error {
	s.Session.Close()
	return nil
}
