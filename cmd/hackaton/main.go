package main

import (
	"fmt"
	"hackathon/backend/config"
	"hackathon/backend/internal/delivery"
	"hackathon/backend/internal/repository"
	"hackathon/backend/internal/service"
	_ "hackathon/backend/pkg/migrate"
	"hackathon/backend/pkg/postgres"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	conf, err := config.New()
	if err != nil {
		logger.Sugar().Fatalf("failed to read new config: %s", err.Error())
	}

	psql, err := postgres.New(&conf.Databases)
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to postgres: %s", err.Error())
	}

	defer psql.Close()

	repo := repository.New(psql)
	svc := service.New(repo)
	app := delivery.New(svc)
	app.Run(fmt.Sprintf(":%d", conf.Application.Port))
}
