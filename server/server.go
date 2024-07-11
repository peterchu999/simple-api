package server

import (
	"log"
	controller "peterchu999/simple-api/controller"
	"peterchu999/simple-api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupServer() *gin.Engine {
	// engine of gin
	loadEnvErr := godotenv.Load()

	if loadEnvErr != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	// database setup
	model.ConnectDatabase()

	// routes
	r.GET("/books", controller.FindBook)
	r.GET("/books/:id", controller.FindBookById)
	r.POST("/books", controller.CreateBook)
	r.PATCH("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)

	return r
}
