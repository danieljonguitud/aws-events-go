package controllers

import (
	"encoding/json"
	"net/http"

	"danieljonguitud.com/aws-events-go/db"
)

type Controller struct {
	Queries *db.Queries
}

func New(queries *db.Queries) *Controller {
	return &Controller{
		Queries: queries,
	}
}

func ResponseHandler(w http.ResponseWriter, status int, response map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
