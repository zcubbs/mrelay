package db

import (
	"github.com/uptrace/bun"
	"github.com/zcubbs/mrelay/cmd/server/config"
)

type MailStore interface {
	// Define CRUD operations for the Mail struct here
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
