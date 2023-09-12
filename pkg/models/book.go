package models

import (
	"github.com/aksentijevicd1/go-mysql-project/pkg/config"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Author      string `gorm:"column:author" json:"author"`
	Name        string `gorm:"column:name" json:"name"`
	Publication string `gorm:"column:publication" json:"publication"`
}

var db *gorm.DB

type Books []*Book

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetBooks() []Book {
	var newBook []Book
	db.Find(&newBook)
	return newBook
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBookById(id int) (*Book, *gorm.DB) {
	var newBook Book
	db := db.Where("ID=?", id).Find(&newBook)
	return &newBook, db
}

func DeleteBook(Id int) Book {
	var newBook Book
	db.Where("ID=?", Id).Delete(&newBook)
	return newBook
}
