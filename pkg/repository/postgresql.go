package repository

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	migrator "github.com/qazaqpyn/api-notz/internal/db"
	"github.com/sirupsen/logrus"
)

func NewPostgresDB(url string) (*sqlx.DB, error) {
	// Recover Migrator
	migrator := migrator.MustGetNewMigrator()

	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	defer db.Close()

	// Apply migrations
	err = migrator.ApplyMigrations(db.DB)
	if err != nil {
		panic(err)
	}

	logrus.Print("Migrations applied")

	return db, nil
}
