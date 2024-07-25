package routers

import (
	"mux-demo/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/book", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/api/book/{id}", handlers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/api/book/{id}", handlers.UpdateBook).Methods("PUT")

	return router
}
