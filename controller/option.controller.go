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

func DeleteNextRule(rw http.ResponseWriter, req *http.Request) {
	var nextRule model.OptionNextRule

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&nextRule); err != nil {
		panic(err)
		return
	}

	db.Where("rule_id", nextRule.RuleID).Where("option_id", nextRule.OptionID).Delete(&nextRule)

	log.Printf("Next Rule [%s] was Deleted...\n", nextRule.RuleID)
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Next Rule [%s] was Deleted!", nextRule.RuleID),
	})
}

func UpdateNextRule(rw http.ResponseWriter, req *http.Request) {
	var nextRule model.OptionNextRule

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&nextRule); err != nil {
		panic(err)
		return
	}

	db.Model(&nextRule).Where("option_id = ?", nextRule.OptionID).Update("rule_id", nextRule.RuleID)

	log.Printf("Rule flow from option \"%v\" to \"%s\" was updated...\n", nextRule.OptionID, nextRule.RuleID)
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Rule flow from option \"%v\" to \"%s\" was updated!", nextRule.OptionID, nextRule.RuleID),
	})
}

// InsertNextRule Insert Next Rule
func InsertNextRule(rw http.ResponseWriter, req *http.Request) {
	var nextRule model.OptionNextRule

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&nextRule); err != nil {
		panic(err)
		return
	}

	db.Create(&nextRule)

	log.Printf("Rule flow from option \"%v\" to \"%s\" was created...", nextRule.OptionID, nextRule.RuleID)
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Rule flow from option \"%v\" to \"%s\" was created!", nextRule.OptionID, nextRule.RuleID),
	})
}

func DeleteOption(rw http.ResponseWriter, req *http.Request) {
	db.Delete(&model.Option{}, mux.Vars(req)["id"])

	log.Printf("Option \"%v\" was deleted...\n", mux.Vars(req)["id"])
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Option \"%v\" was deleted!", mux.Vars(req)["id"]),
	})
}

func UpdateOption(rw http.ResponseWriter, req *http.Request) {
	var option model.Option

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&option); err != nil {
		panic(err)
		return
	}

	db.Model(&option).Where("id", option.ID).Updates(map[string]interface{}{"description": option.Description, "rule_id": option.RuleID})

	log.Printf("Option \"%v\" was updated...\n", option.ID)
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Option \"%v\" was updated!", option.ID),
	})
}

// InsertOption Insert an Option
func InsertOption(rw http.ResponseWriter, req *http.Request) {
	var option model.Option

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&option); err != nil {
		panic(err)
		return
	}

	db.Create(&option)

	log.Println("Option was registered...")
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: fmt.Sprintf("Option \"%s\" was registered for rule \"%s\"!", option.Description, option.RuleID),
	})
}
