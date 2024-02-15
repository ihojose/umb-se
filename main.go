package main

import (
	"airbusexpert/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("==> Starting UMB - Expert System...")
	log.Println("==> Airbus Aircraft Diagnosis by Jose Buelvas Santos.")
	log.Println("==> VERSION 1.0.2")

	// Api Init
	router := mux.NewRouter()
	api.Router(router)
	http.Handle("/", router)

	// Handlers
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"*"})

	// Server Init
	log.Println("Connected to port 5000.")
	log.Fatalln(http.ListenAndServe(":5000", handlers.CORS(credentials, methods, origins, headers)(router)))
}
