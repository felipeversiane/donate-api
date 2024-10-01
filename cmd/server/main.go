package main

import (
	"context"

	"github.com/felipeversiane/donate-api/internal/infra/config"
	"github.com/felipeversiane/donate-api/internal/infra/config/log"
	"github.com/felipeversiane/donate-api/internal/infra/server"
	"github.com/felipeversiane/donate-api/internal/infra/services/cloud"
	"github.com/felipeversiane/donate-api/internal/infra/services/database"
)

func main() {
	cfg := config.NewConfig()

	logger := log.NewLogger(cfg.GetLogConfig())
	logger.Configure()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database := database.NewDatabaseConnection(ctx, cfg.GetDatabaseConfig())
	defer database.Close()

	cloudService := cloud.NewCloudService(cfg.GetCloudServiceConfig())
	objectStorage := cloud.NewObjectStorage(cloudService.GetSession(), cfg.GetCloudServiceConfig())
	objectStorage.CreateBucket(context.Context(context.TODO()))

	server := server.NewServer(cfg.GetServerConfig(), database, objectStorage)
	server.SetupRouter()
	server.Start()
}
