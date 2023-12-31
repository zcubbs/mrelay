package db

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/zcubbs/mrelay/cmd/server/config"
	"github.com/zcubbs/mrelay/cmd/server/models"
)

type MailStore interface {
	SaveMail(mail *models.Email) error
}

func NewMailStore(conn *bun.DB, cfg config.DatabaseConfig) MailStore {
	if cfg.Sqlite.Enabled {
		return &SqliteMailStore{
			conn: conn,
		}
	} else if cfg.Postgres.Enabled {
		return &PostgresMailStore{
			conn: conn,
		}
	}
	return nil
}

type PostgresMailStore struct {
	conn *bun.DB
}

type SqliteMailStore struct {
	conn *bun.DB
}

func (s *PostgresMailStore) SaveMail(mail *models.Email) error {
	_, err := s.conn.NewInsert().Model(mail).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (s *SqliteMailStore) SaveMail(mail *models.Email) error {
	_, err := s.conn.NewInsert().Model(mail).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
