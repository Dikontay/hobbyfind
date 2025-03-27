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
	queryBuilder := qb.Select("users").Columns("id", "name", "email", "fullname", "phone", "password", "created_at")

	dbSession := s.session.GetSession()
	// Динамически добавляем условия фильтрации
	if user.ID != "" {
		queryBuilder = queryBuilder.Where(qb.Eq("id"))
	}
	if user.Username != "" {
		queryBuilder = queryBuilder.Where(qb.Eq("name"))
	}
	if user.Email != "" {
		queryBuilder = queryBuilder.Where(qb.Eq("email"))
	}
	if user.Fullname != "" {
		queryBuilder = queryBuilder.Where(qb.Eq("fullname"))
	}
	if user.Phone != "" {
		queryBuilder = queryBuilder.Where(qb.Eq("phone"))
	}

	// Генерируем CQL-запрос
	stmt, names := queryBuilder.ToCql()

	// Создаем слайс для хранения результатов
	var users []*models.User

	// Создаем новый запрос с привязкой переданных параметров
	query := dbSession.Query(stmt, names)
	bindParams := []interface{}{}

	if user.ID != "" {
		bindParams = append(bindParams, user.ID)
	}
	if user.Username != "" {
		bindParams = append(bindParams, user.Username)
	}
	if user.Email != "" {
		bindParams = append(bindParams, user.Email)
	}
	if user.Fullname != "" {
		bindParams = append(bindParams, user.Fullname)
	}
	if user.Phone != "" {
		bindParams = append(bindParams, user.Phone)
	}

	// Привязываем параметры к запросу
	query.Bind(bindParams...)

	// Выполняем запрос
	iter := query.Iter()

	// Итерация по результатам запроса
	var copyUser models.User
	for iter.Scan(&copyUser.ID,
		&copyUser.Username,
		&copyUser.Email,
		&copyUser.Fullname,
		&copyUser.Phone,
		&copyUser.Password,
		&copyUser.CreatedAt) {
		// Создаем копию объекта пользователя, чтобы избежать перезаписи данных
		u := copyUser
		users = append(users, &u)
	}

	// Проверяем ошибки итерации
	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}

	return users, nil
}

func (s service) Stop() {
	//TODO implement me
	panic("implement me")
}
