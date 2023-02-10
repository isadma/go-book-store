package main

import (
	"fmt"
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
	fmt.Println("Server is started on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
