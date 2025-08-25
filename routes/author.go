package routes

import (
	"golang-restApi/controllers"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", controllers.IndexAuthor).Methods("GET")
	router.HandleFunc("", controllers.CreateAuthor).Methods("POST")
	router.HandleFunc("/{id}", controllers.DetailAuthor).Methods("GET")
	router.HandleFunc("/{id}", controllers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteAuthor).Methods("DELETE")
}
