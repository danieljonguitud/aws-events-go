package controllers

import (
	"encoding/json"
	"net/http"
)

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
	})
}
