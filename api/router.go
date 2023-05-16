package api

import (
	"airbusexpert/controller"
	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	router.HandleFunc("/rule", controller.InsertRule).Methods("POST")
	router.HandleFunc("/rule", controller.UpdateRule).Methods("PUT")
	router.HandleFunc("/rule/{id}", controller.DeleteRule).Methods("DELETE")
	router.HandleFunc("/rule/{id}", controller.GetRule).Methods("GET")
	router.HandleFunc("/rule", controller.GetRules).Methods("GET")
	router.HandleFunc("/rule/next", controller.InsertNextRule).Methods("POST")
	router.HandleFunc("/rule/next", controller.UpdateNextRule).Methods("PUT")
	router.HandleFunc("/rule/next", controller.DeleteNextRule).Methods("DELETE")
	router.HandleFunc("/option", controller.InsertOption).Methods("POST")
	router.HandleFunc("/option", controller.UpdateOption).Methods("PUT")
	router.HandleFunc("/option/{id}", controller.DeleteOption).Methods("DELETE")
	router.HandleFunc("/account", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/account", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/account/login", controller.UserLogin).Methods("POST")
}
