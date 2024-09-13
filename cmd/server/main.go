package main

import (
	"context"

	"github.com/felipeversiane/donate-api/internal/infra/config"
	"github.com/felipeversiane/donate-api/internal/infra/database"
	"github.com/felipeversiane/donate-api/internal/infra/log"
	"github.com/felipeversiane/donate-api/internal/infra/server"
)

func main() {
	cfg := config.NewConfig()

	logger := log.NewLogger(cfg.GetLogConfig())
	logger.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database := database.NewDatabaseConnection(ctx, cfg.GetDatabaseConfig())
	defer database.Close()

	server := server.NewServer(cfg.GetServerConfig(), database)
	server.SetupRouter()
	server.Start()
}
