package cassandra

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
	"github.com/Dikontay/hobbyfind/repository/internal/services/cassandra/migration"
	"github.com/Dikontay/hobbyfind/repository/internal/services/cassandra/session"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	"time"
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
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	user.ID = id.String()

	dbSession := s.session.GetSession()
	stmt, names := qb.Insert("users").Columns(
		"id", "name", "email", "fullname", "phone", "password", "created_at").ToCql()

	user.CreatedAt = new(time.Time)
	*user.CreatedAt = time.Now()

	query := dbSession.
		Query(stmt, names).
		Bind(user.ID, user.Username, user.Email, user.Fullname, user.Phone, user.Password, user.CreatedAt)

	if err = query.Exec(); err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	return &user, nil
}

func (s service) ReadUser(id string) (*models.User, error) {

	dbSession := s.session.GetSession()
	stmt, names := qb.Select("users").Where(qb.Eq("id")).ToCql()

	user := models.User{}
	query := dbSession.Query(stmt, names).Bind(id)
	err := query.Scan(&user)
	if err != nil {
		return &models.User{}, fmt.Errorf("failed to get user: %v", err)
	}

	return &user, nil
}

func (s service) UpdateUser(user models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}

func (s service) ListUsers(user models.User) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Stop() {
	//TODO implement me
	panic("implement me")
}
