package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Threx-code/go-bookstore/package/models"
	"github.com/Threx-code/go-bookstore/package/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Books

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	response, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetABook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// the model.GetABook is returning two variables the book and the DB. Becuase we dont want to use the DB we replace it with underscore _
	bookDetails, _ := models.GetABook(ID)
	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Books{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	response, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Books{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, db := models.GetABook(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(bookDetails)

	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
