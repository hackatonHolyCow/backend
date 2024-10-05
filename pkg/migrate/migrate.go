package migrate

import (
	"database/sql"
	"hackathon/backend/config"
	_ "hackathon/backend/pkg/postgres/migrations"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func init() {
	log := zap.Must(zap.NewProduction())

	conf, err := config.New()
	if err != nil {
		log.Sugar().Fatalf("config: failed to instance new config: %v", err)
	}

	db, err := sql.Open("postgres", conf.Databases.PostgresDSN)
	if err != nil {
		log.Sugar().Fatalf("goose: failed to open postgres: %v", err)
	}

	defer db.Close()

	if err := goose.Up(db, "./pkg/postgres/migrations"); err != nil {
		log.Sugar().Fatalf("goose: failed to up migrations: %v", err)
	}
}
