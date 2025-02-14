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

	err = s.migration.Start(s.session.GetSession(), s.configs.Session.KeySpace)
	if err != nil {
		return fmt.Errorf("failed to init migration: %v", err)
	}
	return nil
}

func (s service) CreateUser(user models.User) (*models.User, error) {
	id, err := gocql.RandomUUID()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate UUID: %v", err)
	}

	user.ID = id.String()

	dbSession := s.session.GetSession()
	stmt, names := qb.Insert("users").Columns(
		"id", "name", "email", "fullname", "phone", "password").ToCql()

	if err = dbSession.Query(stmt, names).Bind(user).Exec(); err != nil {
		return nil, fmt.Errorf("Failed to insert user: %v", err)
	}

	return &user, nil
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
