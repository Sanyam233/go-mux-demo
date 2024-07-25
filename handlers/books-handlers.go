package handlers

import (
	"encoding/json"
	"fmt"
	"mux-demo/helpers"
	"mux-demo/models"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	1. How to pass a http status code to a request
	2. error handling in go
	3. awaiting for api requests in go
*/

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(helpers.CreateResponse(models.Books, []interface{}{}))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var result interface{}
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// models.Books = append(models.Books, book)

	// json.NewEncoder(w).Encode(helpers.CreateResponse([]models.Book{book}, []interface{}{}))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := mux.Vars(r)

	for _, item := range models.Books {
		fmt.Printf("%v, %v\n", item.ID, queryParams["id"])
		if item.ID == queryParams["id"] {
			json.NewEncoder(w).Encode(helpers.CreateResponse([]models.Book{item}, []interface{}{}))
			return
		}
	}

	json.NewEncoder(w).Encode(helpers.CreateResponse([]interface{}{}, []interface{}{}))

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
