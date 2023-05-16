package main

import (
	"airbusexpert/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("==> Starting UMB - Expert System...")
	log.Println("==> Airbus Aircraft Diagnosis by Jose Buelvas Santos.")
	log.Println("==> VERSION 1.0.1")

	// Api Init
	router := mux.NewRouter()
	api.Router(router)
	http.Handle("/", router)

	// Server Init
	log.Println("Connected to port 1234.")
	log.Fatalln(http.ListenAndServe(":1234", router))
}
