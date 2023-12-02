package db

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
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

	// perform migrations
	err = RunMigrations(db)
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	// connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func initSqlite(cfg config.SQLiteConfig) (*bun.DB, error) {
	dsn := cfg.Datasource // ex: database.db
	sqlDb, err := sql.Open(sqliteshim.ShimName, dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite database: %w", err)
	}

	// Create a new bun.DB instance
	dialect := sqlitedialect.New()
	db := bun.NewDB(sqlDb, dialect)
	return db, nil
}

func initPostgres(cfg config.PostgresConfig) (*bun.DB, error) {
	// Construct the Postgres DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// Open the Postgres database
	sqlDb, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening postgres database: %w", err)
	}

	// Create a new bun.DB instance
	db := bun.NewDB(sqlDb, pgdialect.New())
	return db, nil
}
