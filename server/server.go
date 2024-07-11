package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	// engine of gin
	r := gin.Default()

	// routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello 111"})
	})

	return r
}
