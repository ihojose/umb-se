package controller

import (
	"airbusexpert/model"
	"airbusexpert/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func DeleteRule(rw http.ResponseWriter, req *http.Request) {
	db.Delete(&model.Rule{}, mux.Vars(req)["id"])

	log.Printf("Rule \"%s\" was deleted...\n", mux.Vars(req)["id"])
	utils.SetResponse(req, rw, model.Response{
		Status:  http.StatusOK,
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

	db.Model(&rule).Where("id", rule.ID).Update("rule", rule.Rule)

	log.Printf("Rule \"%s\" was updated...\n", rule.ID)
	utils.SetResponse(req, rw, model.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Rule \"%s\" was updated!", rule.ID),
	})
}

// GetRule Get One Rule...
func GetRule(rw http.ResponseWriter, req *http.Request) {
	var rule model.Rule

	res := db.First(&rule, "id", mux.Vars(req)["id"])
	db.Model(&model.Rule{}).Preload("Options").Preload("Options.NextRule").Preload("Options.HasProblem").Preload("Options.HasProblem.Problem").Preload("Options.HasProblem.Problem.Solution").Find(&rule, "id", mux.Vars(req)["id"])

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		utils.SetResponseCode(req, rw, model.Response{
			Status:  http.StatusNotFound,
			Message: "Ups! Rule not found!",
		})
		return
	}

	utils.SetResponse(req, rw, rule)
}

// GetRules Get All Rules...
func GetRules(rw http.ResponseWriter, req *http.Request) {
	var rules []model.Rule

	db.Model(&[]model.Rule{}).Preload("Options").Preload("Options.NextRule").Preload("Options.HasProblem.Problem").Find(&rules)

	utils.SetResponse(req, rw, rules)
}

// InsertRule Insert One Rule...
func InsertRule(rw http.ResponseWriter, req *http.Request) {
	var rule model.Rule

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&rule); err != nil {
		panic(err)
		return
	}

	db.Create(&rule)

	log.Printf("Rule \"%s\" was created...\n", rule.ID)
	utils.SetResponse(req, rw, model.Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Rule \"%s\" was created successfully!", rule.ID),
	})
}
