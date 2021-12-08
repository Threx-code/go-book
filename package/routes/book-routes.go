package routes

import (
	"github.com/Threx-code/go-bookstore/package/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.GetABook).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
