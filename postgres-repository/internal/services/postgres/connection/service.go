package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"time"
)

type service struct {
	configs Configs
	bunDb   *bun.DB
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
		bunDb:   nil,
	}
}

func (s *service) GetClient(args ...int) (*bun.DB, error) {
	if s.bunDb != nil {
		if err := s.bunDb.Ping(); err != nil {
			count := 5
			if len(args) > 0 {
				count = args[0]
			}
			if count == 0 {
				return nil, fmt.Errorf("failed to connect to database, not retry left")
			}

			s.bunDb.Close()
			s.bunDb = nil
			return s.GetClient(count - 1)
		} else {
			return s.bunDb, nil
		}
	}

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		s.configs.User,
		s.configs.Password,
		s.configs.Host,
		s.configs.Port,
		s.configs.DB,
		s.configs.SslMode,
		s.configs.Schema,
	)

	sqlDb, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	bunDb := bun.NewDB(sqlDb, pgdialect.New())

	err = bunDb.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	bunDb.SetMaxIdleConns(3)
	bunDb.SetConnMaxIdleTime(60 * time.Second)

	s.bunDb = bunDb

	return bunDb, nil
}
