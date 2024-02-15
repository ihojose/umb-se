package utils

import (
	"airbusexpert/model"
	"encoding/json"
	"net/http"
)

// SetResponse Set Api Json Response
func SetResponse(r *http.Request, w http.ResponseWriter, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

// SetResponseCode SetResponse Set Api Json Response
func SetResponseCode(r *http.Request, w http.ResponseWriter, v model.Response) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(v.Status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}
