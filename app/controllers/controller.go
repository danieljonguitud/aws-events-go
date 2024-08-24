package controllers

import (
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

func RenderView(w http.ResponseWriter, status int, response map[string]interface{}) {
}
