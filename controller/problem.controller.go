package controller

import (
	"airbusexpert/model"
	"airbusexpert/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteHasProblem(rw http.ResponseWriter, req *http.Request) {
	db.Delete(&model.OptionHasProblem{}).Where("option_id", mux.Vars(req)["id"]).Where("problem_id", mux.Vars(req)["pr"])

	log.Printf("Has Problem \"%s\" was deleted...\n", mux.Vars(req)["id"])
	utils.SetResponse(req, rw, model.Response{
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

	db.Model(&problem).Where("option_id", problem.ProblemID).Updates(map[string]interface{}{
		"problem_id": problem.ProblemID,
	})

	log.Printf("Has problem \"%v\" was updated...\n", problem.ProblemID)
	utils.SetResponse(req, rw, model.Response{
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

	db.Create(&has)

	log.Println("Option Has Problem was registered...")
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Option Has Problem was registered!"),
	})
}

func HasProblem(rw http.ResponseWriter, req *http.Request) {
	var has []model.OptionHasProblem

	db.Find(&has, "option_id", mux.Vars(req)["id"])

	utils.SetResponse(req, rw, has)
}

func GetProblems(rw http.ResponseWriter, req *http.Request) {
	var problems []model.Problem

	db.Find(&problems)

	utils.SetResponse(req, rw, problems)
}

func GetProblem(rw http.ResponseWriter, req *http.Request) {
	var problem model.Problem

	db.Model(&model.Problem{}).Preload("Solution").Find(&problem, "id", mux.Vars(req)["id"])

	utils.SetResponse(req, rw, problem)
}

func DeleteProblem(rw http.ResponseWriter, req *http.Request) {
	db.Delete(&model.Problem{}, mux.Vars(req)["id"])

	log.Printf("Problem \"%v\" was deleted...\n", mux.Vars(req)["id"])
	utils.SetResponse(req, rw, model.Response{
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

	db.Model(&problem).Where("id", problem.ID).Updates(map[string]interface{}{
		"description":   problem.Description,
		"urgency_level": problem.UrgencyLevel,
	})

	log.Printf("Problem \"%v\" was updated...\n", problem.ID)
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Problem \"%v\" was updated!", problem.ID),
	})
}

func InsertProblem(rw http.ResponseWriter, req *http.Request) {
	var problem model.Problem
	//var solution model.Solution

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&problem); err != nil {
		panic(err)
		return
	}

	// Save problem
	if err := db.Create(&problem).Error; err != nil {
		panic(err)
		return
	}

	// Get problem
	db.First(&problem, "id", problem.ID)

	// Save Option has Problem
	ophp := &model.OptionHasProblem{
		ProblemID: problem.ID,
		OptionID:  problem.OptionId,
	}
	if err := db.Create(&ophp).Error; err != nil {
		panic(err)
		return
	}

	log.Println("Problem was registered...")
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Problem \"%s\" was registered!", problem.Description),
	})
}
