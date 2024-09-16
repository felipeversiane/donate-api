package server

import (
	"fmt"
	"log/slog"

	"github.com/felipeversiane/donate-api/internal/infra/config"
	"github.com/felipeversiane/donate-api/internal/infra/services/aws"
	"github.com/felipeversiane/donate-api/internal/infra/services/database"
	"github.com/gin-gonic/gin"
)

type ServerInterface interface {
	SetupRouter()
	Start()
}

type server struct {
	router        *gin.Engine
	config        config.ServerConfig
	database      database.DatabaseInterface
	objectStorage aws.ObjectStorageInterface
}

func NewServer(
	cfg config.ServerConfig,
	database database.DatabaseInterface,
	objectStorage aws.ObjectStorageInterface,
) ServerInterface {
	server := &server{
		router:        gin.New(),
		config:        cfg,
		database:      database,
		objectStorage: objectStorage,
	}
	return server
}

func (s *server) SetupRouter() {
	s.router.Use(gin.Recovery())
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
}

func (s *server) Start() {
	port := s.config.Port
	if port == "" {
		port = "8000"
	}
	err := s.router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
	}
}
