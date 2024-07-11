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
	r.GET("/books/:id", controller.FindBookById)
	r.POST("/books", controller.CreateBook)
	r.PATCH("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)

	return r
}
