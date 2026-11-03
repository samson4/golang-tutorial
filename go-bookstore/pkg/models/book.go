package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/samson4/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) *Book {
	var BookById Book
	result := db.Where("ID=?", id).Find(&BookById)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return &BookById
}

func DeleteBookById(id int64) *Book {
	var BookById Book
	result := db.Where("ID=?", id).Delete(BookById)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return &BookById
}
