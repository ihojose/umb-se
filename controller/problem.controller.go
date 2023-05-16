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

func DeleteHasProblem(rw http.ResponseWriter, req *http.Request) {
	db := database.Connect()
	db.Delete(&model.OptionHasProblem{}).Where("option_id", mux.Vars(req)["id"]).Where("problem_id", mux.Vars(req)["pr"])

	log.Printf("Has Problem \"%s\" was deleted...\n", mux.Vars(req)["id"])
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Has Problem \"%s\" was deleted!", mux.Vars(req)["id"]),
	})
}

func UpdateHasProblem(rw http.ResponseWriter, req *http.Request) {
	var problem model.OptionHasProblem

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&problem); err != nil {
		panic(err)
		return
	}

	db := database.Connect()
	db.Model(&problem).Where("option_id", problem.ProblemID).Updates(map[string]interface{}{
		"problem_id": problem.ProblemID,
	})

	log.Printf("Has problem \"%v\" was updated...\n", problem.ProblemID)
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Has problem \"%v\" was updated!", problem.ProblemID),
	})
}

func InsertHasProblem(rw http.ResponseWriter, req *http.Request) {
	var has model.OptionHasProblem

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&has); err != nil {
		panic(err)
		return
	}

	db := database.Connect()
	db.Create(&has)

	log.Println("Option Has Problem was registered...")
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Option Has Problem was registered!"),
	})
}

func HasProblem(rw http.ResponseWriter, req *http.Request) {
	var has []model.OptionHasProblem

	db := database.Connect()
	db.Find(&has, "option_id", mux.Vars(req)["id"])

	utils.SetResponse(rw, has)
}

func GetProblems(rw http.ResponseWriter, req *http.Request) {
	var problems []model.Problem

	db := database.Connect()
	db.Find(&problems)

	utils.SetResponse(rw, problems)
}

func GetProblem(rw http.ResponseWriter, req *http.Request) {
	var problem model.Problem

	db := database.Connect()
	db.Model(&model.Problem{}).Preload("Solution").Find(&problem, "id", mux.Vars(req)["id"])

	utils.SetResponse(rw, problem)
}

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
