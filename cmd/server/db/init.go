package db

import (
	"fmt"
	"github.com/uptrace/bun"
	"github.com/zcubbs/mrelay/cmd/server/config"
)

func InitializeDB(cfg config.DatabaseConfig) (*bun.DB, error) {
	var db *bun.DB
	var err error

	if cfg.Sqlite.Enabled && cfg.Postgres.Enabled {
		return nil, fmt.Errorf("both sqlite and postgres are enabled, please only enable one")
	}

	if cfg.Sqlite.Enabled {
		db, err = initSqlite(cfg.Sqlite)
		if err != nil {
			return nil, fmt.Errorf("error initializing sqlite: %w", err)
		}
	} else if cfg.Postgres.Enabled {
		db, err = initPostgres(cfg.Postgres)
		if err != nil {
			return nil, fmt.Errorf("error initializing postgres: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no database enabled")
	}

	return db, nil
}

func initSqlite(cfg config.SQLiteConfig) (*bun.DB, error) {
	return nil, nil
}

func initPostgres(cfg config.PostgresConfig) (*bun.DB, error) {
	return nil, nil
}
