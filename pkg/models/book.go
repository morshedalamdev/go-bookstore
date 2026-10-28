package models

import (
	"go-bookstore/pkg/config"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Author      string `json:"author" gorm:"type:varchar(255);not null"`
	Publication string `json:"publication" gorm:"type:varchar(255);not null"`
}

var db = config.GetDB()

// create a new book record
func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

// get all books records
func GetAllBooks() ([]Book, error) {
	var books []Book

	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

// get a book record by ID
func GetBookById(id int64) (*Book, error) {
	var book Book

	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

// delete a book record by ID
func DeleteBookById(id int64) (*Book, error) {
	var book Book

	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
