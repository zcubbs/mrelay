package repository

import (
	"github.com/uptrace/bun"
	"github.com/zcubbs/mrelay/cmd/server/config"
)

type MailRepository interface {
	// Define CRUD operations for the Mail struct here
}

func NewMailRepository(conn *bun.DB, cfg config.DatabaseConfig) MailRepository {
	if cfg.Sqlite.Enabled {
		return &SqliteMailRepository{
			conn: conn,
		}
	} else if cfg.Postgres.Enabled {
		return &PostgresMailRepository{
			conn: conn,
		}
	}
	return nil
}

type SqliteMailRepository struct {
	conn *bun.DB
}

type PostgresMailRepository struct {
	conn *bun.DB
}
