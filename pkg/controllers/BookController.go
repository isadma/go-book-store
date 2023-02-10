package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/isadma/go-book-store/pkg/models"
	"github.com/isadma/go-book-store/pkg/utils"
)

var NewBook models.Book

func Index(response http.ResponseWriter, request *http.Request) {
	books := models.Index()
	res, _ := json.Marshal(books)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func Show(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID on show")
	}
	book, _ := models.Show(Id)
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func Create(response http.ResponseWriter, request *http.Request) {
	model := &models.Book{}
	utils.ParseBody(request, model)
	book := model.Create()
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func Update(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID on update")
	}
	book, db := models.Show(Id)

	model := &models.Book{}
	utils.ParseBody(request, model)

	book.Name = model.Name
	book.Author = model.Author
	book.Publication = model.Publication

	db.Save(&book)
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID on delete")
	}
	book := models.Delete(Id)
	res, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
