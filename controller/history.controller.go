package controller

import (
	"airbusexpert/model"
	"airbusexpert/utils"
	"airbusexpert/utils/dates"
	"encoding/json"
	"net/http"
	"time"
)

func AddHistory(rw http.ResponseWriter, req *http.Request) {
	var history model.History

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&history); err != nil {
		panic(err)
		return
	}

	// Complete data
	history.Datetime = time.Now().Format(dates.DateFormat("YYYYMMDDHHmmss"))

	// Save
	if err := db.Create(&history).Error; err != nil {
		panic(err)
		return
	}

	utils.SetResponse(req, rw, model.Response{
		Status:  200,
		Message: "History Saved",
	})
}
