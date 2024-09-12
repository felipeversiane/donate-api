package main

import (
	"github.com/felipeversiane/donate-api/internal/infra/api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	g.Use(gin.Recovery())
	router.SetupRoutes(g)
	g.Run()
}
