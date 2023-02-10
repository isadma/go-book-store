package routes

import (
	"github.com/gorilla/mux"
	"github.com/isadma/go-book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.Index).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.Show).Methods("GET")
	router.HandleFunc("/books", controllers.Create).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.Delete).Methods("DELETE")
}
