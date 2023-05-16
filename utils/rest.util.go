package utils

import (
	"encoding/json"
	"net/http"
)

// SetResponse Set Api Json Response
func SetResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}
