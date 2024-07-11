package repository

import (
	. "peterchu999/simple-api/dto"
	"peterchu999/simple-api/model"
)

func GetAll() []model.Book {
	// create a new variable with type of []model.Book
	var books []model.Book
	// passing books variable, to be populate by the DB
	// model.DB.Model(&model.Book{}).Find(&books)
	//short hand
	model.DB.Find(&books)

	return books
}

func GetById(id int) (model.Book, error) {
	var book model.Book
	if err := model.DB.First(&book, "id = ?", id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func Create(bookDto CreateBookDto) (model.Book, error) {
	book := model.Book{Title: bookDto.Title, Author: bookDto.Author}

	if err := model.DB.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func Update(id int, bookDto UpdateBookDto) (model.Book, error) {
	book := model.Book{}

	res := model.DB.Model(book).Where("id=?", id).Updates(bookDto)
	res.First(&book)
	if res.Error != nil {
		return book, res.Error
	}

	return book, nil
}

func Delete(id int) (bool, error) {
	res := model.DB.Where("id=?", id).Delete(&model.Book{})
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected >= 1, nil
}
