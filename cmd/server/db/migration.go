package db

import (
	"context"
	"embed"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/zcubbs/log"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func RunMigrations(db *bun.DB) error {
	var migrations = migrate.NewMigrations()

	if err := migrations.Discover(migrationsFS); err != nil {
		return err
	}

	ctx := context.Background()
	// Create a new Migrator instance
	migrator := migrate.NewMigrator(db, migrations)

	// Init migrations table
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	// Apply the migrations
	mg, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if mg.IsZero() {
		log.Info("no database migrations to run")
		return nil
	}

	log.Info("migrations ran successfully",
		"migrations", mg.Migrations.String(),
	)
	return nil
}
