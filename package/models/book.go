package models

import (
	"github.com/Threx-code/go-bookstore/package/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Books struct {
	// gorm.Model
	gorm.Model
	Name        string `gorm: "" json:name"`
	Author      string `json:author`
	Publication string `json: "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Books{})
}

func (b *Books) CreateBook() *Books {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Books {
	var AllBooks []Books
	db.Find(&AllBooks)
	return AllBooks
}

func GetABook(Id int64) (*Books, *gorm.DB) {
	var getBook Books
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Books {
	var book Books
	db.Where("ID=?", ID).Delete(book)
	return book
}
