package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	"project/internal/domain/models"
	"project/internal/services/cassandra/migration"
	"project/internal/services/cassandra/session"
)

type service struct {
	configs   Configs
	session   session.Service
	migration migration.Service
}

func NewService(configs Configs) Service {
	return &service{
		configs:   configs,
		session:   session.NewService(configs.Session),
		migration: migration.NewService(configs.Migration),
	}
}

func (s service) Init() error {
	err := s.session.Init()
	if err != nil {
		return fmt.Errorf("failed to init session: %v", err)
	}

	err = s.migration.Start(s.session.GetSession(), s.configs.Session.KeySpace, s.configs.Session.DbName)
	if err != nil {
		return fmt.Errorf("failed to init migration: %v", err)
	}
	return nil
}

func (s service) CreateUser(user models.User) (models.User, error) {
	query := &gocql.Query{}
	qb.Insert(s.configs.Session.DbName).Columns("id", "name", "email", "password").ToCql(query)

}

func (s service) ReadUser(id string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) UpdateUser(user models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}

func (s service) ListUsers(user models.User) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Stop() {
	//TODO implement me
	panic("implement me")
}
