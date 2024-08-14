package controllers

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/db"
)

type Controller struct {
	Queries *db.Queries
	Server  *http.ServeMux
}

func New(queries *db.Queries, server *http.ServeMux) *Controller {
	return &Controller{
		Queries: queries,
		Server:  server,
	}
}

