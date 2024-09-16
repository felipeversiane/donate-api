package main

import (
	"context"

	"github.com/felipeversiane/donate-api/internal/infra/config"
	"github.com/felipeversiane/donate-api/internal/infra/config/log"
	"github.com/felipeversiane/donate-api/internal/infra/server"
	"github.com/felipeversiane/donate-api/internal/infra/services/aws"
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

	cloud := aws.NewCloudService(cfg.GetCloudServiceConfig())
	objectStorage := aws.NewObjectStorage(cloud.GetSession(), cfg.GetCloudServiceConfig())
	objectStorage.CreateBucket(context.Context(context.TODO()))

	server := server.NewServer(cfg.GetServerConfig(), database, objectStorage)
	server.SetupRouter()
	server.Start()
}
