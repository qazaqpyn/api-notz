package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	migrator "github.com/qazaqpyn/api-notz/internal/db"
	"github.com/sirupsen/logrus"
)

func NewPostgresDB(url string) (*sql.DB, error) {
	// Recover Migrator
	migrator := migrator.MustGetNewMigrator()

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
