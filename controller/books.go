package controller

import (
	"log"
	"net/http"
	. "peterchu999/simple-api/dto"
	"peterchu999/simple-api/model"
	bookRepo "peterchu999/simple-api/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindBook(c *gin.Context) {
	var books []model.Book = bookRepo.GetAll()
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func FindBookById(c *gin.Context) {
	// get path parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Println("error while accesing path param", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "not a valid path"})
		return
	}

	book, repoErr := bookRepo.GetById(id)

	if repoErr != nil {
		log.Println("error fetching from DB", repoErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
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

func UpdateBook(c *gin.Context) {

	// get path parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Println("error while accesing path param", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "not a valid path"})
		return
	}

	var input UpdateBookDto
	// get input and validate it to dto
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("error while validating dto", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//update book
	book, err := bookRepo.Update(id, input)
	if err != nil {
		log.Println("error while updating book", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func DeleteBook(c *gin.Context) {
	// get path parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Println("error while accesing path param", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "not a valid path"})
		return
	}

	//update book
	isDeleted, deleteError := bookRepo.Delete(id)
	if deleteError != nil {
		log.Println("error while delete book", deleteError.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": deleteError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": isDeleted})
}
