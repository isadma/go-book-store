package models

import (
	"github.com/isadma/go-book-store/pkg/config"
	"github.com/jinzhu/gorm"
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

func Index() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func Show(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func (book *Book) Create() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func delete(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
