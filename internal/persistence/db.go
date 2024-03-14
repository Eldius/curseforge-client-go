package persistence

import (
	"embed"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"log/slog"
)

var (
	db *sqlx.DB

	//go:embed all:migrations
	migFiles embed.FS
)

func GetDB() (*sqlx.DB, error) {
	if db == nil {
		_db, err := sqlx.Open("sqlite", "mod_cache.db")
		if err != nil {
			err = fmt.Errorf("opening db file: %w", err)
			return nil, err
		}
		db = _db
		if err := runMigrations(); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func runMigrations() error {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migFiles,
		Root:       "migrations",
	}

	migInfo, err := migrations.FindMigrations()
	if err != nil {
		err = fmt.Errorf("finding migrations: %w", err)
		slog.With("error", err).Error("failed to find migrations")
		return err
	}
	log := slog.With(slog.Int("migrations_to_do", len(migInfo)))
	log.Info("running migrations begin")

	n, err := migrate.Exec(db.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		err = errors.Wrap(err, "failed to execute migrations")
		log.With("error", err).Error("failed to run migrations")
		return err
	}
	log.With(slog.Int("migrations_done", n)).Info("migrations end")

	return nil
}
