package main

import (
	"fmt"
	"golang-restApi/config"
	"golang-restApi/routes"
	"net/http"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RoutesIndex(r)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(r)

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), handler)
}
