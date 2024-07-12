package server

import (
	"log"
	controller "peterchu999/simple-api/controller"
	"peterchu999/simple-api/middleware"
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

	r.GET("/login", controller.Login)
	booksRouter := r.Group("/books")
	// routes

	booksRouter.Use(middleware.JWTAuthMiddleware())
	{
		booksRouter.GET("/", controller.FindBook)
		booksRouter.GET("/:id", controller.FindBookById)
		booksRouter.POST("/", controller.CreateBook)
		booksRouter.PATCH("/:id", controller.UpdateBook)
		booksRouter.DELETE("/:id", controller.DeleteBook)
	}

	return r
}
