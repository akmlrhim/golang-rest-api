package routes

import (
	authcontroller "golang-restApi/controllers"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authcontroller.Index).Methods("GET")
	router.HandleFunc("", authcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}", authcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}", authcontroller.Update).Methods("PUT")
}
