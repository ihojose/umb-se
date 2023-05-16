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

func DeleteRule(rw http.ResponseWriter, req *http.Request) {
	db := database.Connect()
	db.Delete(&model.Rule{}, mux.Vars(req)["id"])

	log.Printf("Rule \"%s\" was deleted...\n", mux.Vars(req)["id"])
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Rule \"%s\" was deleted!", mux.Vars(req)["id"]),
	})
}

func UpdateRule(rw http.ResponseWriter, req *http.Request) {
	var rule model.Rule

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&rule); err != nil {
		panic(err)
		return
	}

	db := database.Connect()
	db.Model(&rule).Where("id", rule.ID).Update("rule", rule.Rule)

	log.Printf("Rule \"%s\" was updated...\n", rule.ID)
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Rule \"%s\" was updated!", rule.ID),
	})
}

// GetRule Get One Rule...
func GetRule(rw http.ResponseWriter, req *http.Request) {
	var rule model.Rule

	db := database.Connect()
	db.Model(&model.Rule{}).Preload("Options").Preload("Options.NextRule").Preload("Options.HasProblem").Find(&rule, "id", mux.Vars(req)["id"])

	utils.SetResponse(rw, rule)
}

// GetRules Get All Rules...
func GetRules(rw http.ResponseWriter, req *http.Request) {
	var rules []model.Rule

	db := database.Connect()
	db.Model(&[]model.Rule{}).Preload("Options").Preload("Options.NextRule").Preload("Options.HasProblem.Problem").Find(&rules)

	utils.SetResponse(rw, rules)
}

// InsertRule Insert One Rule...
func InsertRule(rw http.ResponseWriter, req *http.Request) {
	var rule model.Rule

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&rule); err != nil {
		panic(err)
		return
	}

	db := database.Connect()
	db.Create(&rule)

	log.Printf("Rule \"%s\" was created...\n", rule.ID)
	utils.SetResponse(rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Rule \"%s\" was created successfully!", rule.ID),
	})
}
