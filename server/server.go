package server

import (
	controller "peterchu999/simple-api/controller"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	// engine of gin
	r := gin.Default()

	// routes
	r.GET("/books", controller.FindBook)
	r.POST("/books", controller.CreateBook)

	return r
}
