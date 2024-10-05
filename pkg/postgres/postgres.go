package postgres

import (
	"hackathon/backend/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(conf *config.Databases) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", conf.PostgresDSN)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
