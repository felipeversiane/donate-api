package main

import (
	"context"

	"github.com/felipeversiane/donate-api/internal/infra/config"
	"github.com/felipeversiane/donate-api/internal/infra/database"
	"github.com/felipeversiane/donate-api/internal/infra/server"
)

func main() {
	cfg := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database := database.NewDatabaseConnection(ctx, cfg.Database)
	defer database.Close()

	server := server.NewServer(cfg.Server, database)
	server.SetupRouter()
	server.Start()
}
