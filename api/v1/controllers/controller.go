package controllers

import (
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
