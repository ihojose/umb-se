package controller

import (
	"airbusexpert/model"
	"airbusexpert/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func DeleteAirplane(rw http.ResponseWriter, req *http.Request) {
	db.Delete(&model.Airplane{}, "id", mux.Vars(req)["id"])

	log.Println("The airplane was deleted...")
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: "The airplane was deleted!",
	})
}

func UpdateAirplane(rw http.ResponseWriter, req *http.Request) {
	var airplane model.Airplane
	var data model.Airplane

	// Get Json Requests
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&airplane); err != nil {
		panic(err)
		return
	}

	// Database registration
	db.First(&data).Where("id", airplane.ID)

	// User Entity
	data.ModelName = strings.ToUpper(airplane.ModelName)
	data.ModelLine = strings.ToUpper(airplane.ModelLine)
	data.ModelUpdate = strings.ToUpper(airplane.ModelUpdate)
	data.Phabricator = strings.ToUpper(airplane.Phabricator)

	db.Save(&data).Where("id", airplane.ID)

	// Response
	log.Println("The airplane was updated...")
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: "The airplane was updated!",
	})
}

func RegisterAirplane(rw http.ResponseWriter, req *http.Request) {
	var airplane model.Airplane

	// Get Json Requests
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&airplane); err != nil {
		panic(err)
		return
	}

	// Database registration
	db.Create(&airplane)

	// Response
	log.Println("New airplane registered...")
	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: "Airplane registered!",
	})
}
