package controller

import (
	"airbusexpert/model"
	"airbusexpert/utils"
	"airbusexpert/utils/dates"
	"encoding/json"
	"net/http"
	"time"
)

func CreateSession(rw http.ResponseWriter, req *http.Request) {
	var session model.Session

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&session); err != nil {
		panic(err)
		return
	}

	// Complete data
	session.Date = time.Now().Format(dates.DateFormat("YYYYMMDDHHmmss"))

	// Save session
	if err := db.Create(&session).Error; err != nil {
		utils.SetResponseCode(req, rw, model.Response{
			Status:  500,
			Message: err.Error(),
		})
		panic(err)
		return
	}

	// Get Session
	db.First(&session, "user_id", session.UserID)

	// Get session
	utils.SetResponse(req, rw, session)
}
