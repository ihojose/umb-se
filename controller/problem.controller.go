package controller

import (
	"airbusexpert/database"
	"airbusexpert/model"
	"airbusexpert/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteProblem(rw http.ResponseWriter, req *http.Request) {
	db := database.Connect()
	db.Delete(&model.Problem{}, mux.Vars(req)["id"])

	log.Printf("Problem \"%v\" was deleted...\n", mux.Vars(req)["id"])
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Problem \"%v\" was deleted!", mux.Vars(req)["id"]),
	})
}

func UpdateProblem(rw http.ResponseWriter, req *http.Request) {
	var problem model.Problem

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&problem); err != nil {
		panic(err)
		return
	}

	db := database.Connect()
	db.Model(&problem).Where("id", problem.ID).Updates(map[string]interface{}{
		"description":   problem.Description,
		"urgency_level": problem.UrgencyLevel,
	})

	log.Printf("Problem \"%v\" was updated...\n", problem.ID)
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Problem \"%v\" was updated!", problem.ID),
	})
}

func InsertProblem(rw http.ResponseWriter, req *http.Request) {
	var problem model.Problem

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&problem); err != nil {
		panic(err)
		return
	}

	db := database.Connect()
	db.Create(&problem)

	log.Println("Problem was registered...")
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Problem \"%s\" was registered!", problem.Description),
	})
}
