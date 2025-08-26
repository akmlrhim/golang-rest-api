package routes

import (
	"golang-restApi/controllers"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("", controllers.IndexBook).Methods("GET")
	router.HandleFunc("", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/{id}", controllers.DetailBook).Methods("GET")
	router.HandleFunc("/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteBook).Methods("DELETE")
}
