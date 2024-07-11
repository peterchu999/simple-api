package controller

import (
	"log"
	"net/http"
	. "peterchu999/simple-api/dto"
	"peterchu999/simple-api/model"
	bookRepo "peterchu999/simple-api/repository"

	"github.com/gin-gonic/gin"
)

func FindBook(c *gin.Context) {
	var books []model.Book = bookRepo.GetAll()
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func CreateBook(c *gin.Context) {
	var input CreateBookDto
	// get input and validate it to dto
	log.Println("got inside Controller Create Book")

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error while validating dto", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book
	book, err := bookRepo.Create(input)
	if err != nil {
		log.Println("error while creating book", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": book})

}
