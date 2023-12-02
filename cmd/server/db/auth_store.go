package db

import (
	"fmt"
	"github.com/uptrace/bun"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/models"
)

type AuthStore interface {
	GetUserByUsername(username string) (models.User, error)
}

func NewInMemoryAuthStore(accounts map[string]string) AuthStore {
	return &InMemoryAuthStore{accounts}
}

func NewAuthStore(conn *bun.DB, cfg config.DatabaseConfig) AuthStore {
	if cfg.Sqlite.Enabled {
		return &SqliteAuthStore{
			conn: conn,
		}
	} else if cfg.Postgres.Enabled {
		return &PostgresAuthStore{
			conn: conn,
		}
	}
	return nil
}

type PostgresAuthStore struct {
	conn *bun.DB
}

type SqliteAuthStore struct {
	conn *bun.DB
}

type InMemoryAuthStore struct {
	accounts map[string]string
}

func (s *InMemoryAuthStore) GetUserByUsername(username string) (models.User, error) {
	if password, ok := s.accounts[username]; ok {
		return models.User{
			Username: username,
			Password: password,
		}, nil
	}
	return models.User{}, fmt.Errorf("user not found")
}

func (s *PostgresAuthStore) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, fmt.Errorf("not implemented")
}

func (s *SqliteAuthStore) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, fmt.Errorf("not implemented")
}
