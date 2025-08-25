package routes

import (
	"golang-restApi/controllers"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("", controllers.IndexBook).Methods("GET")
	router.HandleFunc("", controllers.CreateBook).Methods("POST")
}
