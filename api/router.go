package api

import (
	"airbusexpert/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(router *mux.Router) {
	router.HandleFunc("/rule", controller.InsertRule).Methods(http.MethodPost)
	router.HandleFunc("/rule", controller.UpdateRule).Methods(http.MethodPut)
	router.HandleFunc("/rule/{id}", controller.DeleteRule).Methods(http.MethodDelete)
	router.HandleFunc("/rule/{id}", controller.GetRule).Methods(http.MethodGet)
	router.HandleFunc("/rule", controller.GetRules).Methods(http.MethodGet)
	router.HandleFunc("/rule/next", controller.InsertNextRule).Methods(http.MethodPost)
	router.HandleFunc("/rule/next", controller.UpdateNextRule).Methods(http.MethodPut)
	router.HandleFunc("/rule/next", controller.DeleteNextRule).Methods(http.MethodDelete)
	router.HandleFunc("/option", controller.InsertOption).Methods(http.MethodPost)
	router.HandleFunc("/option", controller.UpdateOption).Methods(http.MethodPut)
	router.HandleFunc("/option/{id}", controller.DeleteOption).Methods(http.MethodDelete)
	router.HandleFunc("/option/has/problem", controller.InsertHasProblem).Methods(http.MethodPost)
	router.HandleFunc("/option/has/problem", controller.UpdateHasProblem).Methods(http.MethodPut)
	router.HandleFunc("/option/has/problem/{id}", controller.HasProblem).Methods(http.MethodGet)
	router.HandleFunc("/option/has/problem/{id}-{pr}", controller.DeleteHasProblem).Methods(http.MethodDelete)
	router.HandleFunc("/account", controller.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/account", controller.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/account/login", controller.UserLogin).Methods(http.MethodPost)
	router.HandleFunc("/airplane", controller.RegisterAirplane).Methods(http.MethodPost)
	router.HandleFunc("/airplane", controller.UpdateAirplane).Methods(http.MethodPut)
	router.HandleFunc("/airplane/{id}", controller.DeleteAirplane).Methods(http.MethodDelete)
	router.HandleFunc("/problem", controller.InsertProblem).Methods(http.MethodPost)
	router.HandleFunc("/problem", controller.UpdateProblem).Methods(http.MethodPut)
	router.HandleFunc("/problem", controller.GetProblems).Methods(http.MethodGet)
	router.HandleFunc("/problem/{id}", controller.DeleteProblem).Methods(http.MethodDelete)
	router.HandleFunc("/problem/{id}", controller.GetProblem).Methods(http.MethodGet)
	router.HandleFunc("/session", controller.CreateSession).Methods(http.MethodPost)
	router.HandleFunc("/history", controller.AddHistory).Methods(http.MethodPost)
}
