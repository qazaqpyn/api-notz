package repository

import (
	"database/sql"
	"embed"

	_ "github.com/lib/pq"
	"github.com/qazaqpyn/api-notz/internal/db/migrator"
	"github.com/sirupsen/logrus"
)

const migrationsDir = "internal/db/migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

func NewPostgresDB(url string) (*sql.DB, error) {
	// Recover Migrator
	migrator := migrator.MustGetNewMigrator(MigrationsFS, migrationsDir)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	defer db.Close()

	// Apply migrations
	err = migrator.ApplyMigrations(db)
	if err != nil {
		panic(err)
	}

	logrus.Print("Migrations applied")

	return db, nil
}
