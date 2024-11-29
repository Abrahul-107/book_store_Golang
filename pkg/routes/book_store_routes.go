package routes

import (
	"github.com/gorilla/mux"
	"github.com/rahul/go-bookstore/pkg/controllers"
)

// RegisterBookStoreRoutes registers routes for the book store
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")           // Create a new book
	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")              // Get all books
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")   // Get a book by ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")    // Update a book by ID
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") // Delete a book by ID
}
