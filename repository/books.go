package repository

import (
	"log"
	. "peterchu999/simple-api/dto"
	"peterchu999/simple-api/model"
)

func GetAll() []model.Book {
	// create a new variable with type of []model.Book
	var books []model.Book
	// passing books variable, to be populate by the DB
	model.DB.Find(&books)

	return books
}

func Create(bookDto CreateBookDto) (model.Book, error) {
	book := model.Book{Title: bookDto.Title, Author: bookDto.Author}
	log.Println(book)
	if err := model.DB.Create(&book).Error; err != nil {
		return model.Book{}, err
	}

	return book, nil
}
