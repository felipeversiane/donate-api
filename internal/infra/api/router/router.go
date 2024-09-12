package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
