package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isadma/go-book-store/pkg/routes"
)

// import (

// )

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
