package main

import (
	"fmt"
	"log"
	"mux-demo/models"
	"mux-demo/routers"
	"net/http"
)

func main() {
	router := routers.Router()
	fmt.Println("Server running on port 8000")

	author1 := models.CreateAuthorInstance("John", "Doe")
	book1 := models.CreateBookInstance("1", "123", "Book 1", author1)
	models.Books = append(models.Books, *book1)
	book2 := models.CreateBookInstance("2", "234", "Book 2", author1)
	models.Books = append(models.Books, *book2)

	log.Fatal(http.ListenAndServe(":8000", router))
}
