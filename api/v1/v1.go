package v1

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/controllers"
	"danieljonguitud.com/aws-events-go/db"
)

type V1API struct {
	Server     *http.ServeMux
	Controller *controllers.Controller
}

func New(server *http.ServeMux, controller *controllers.Controller, db *db.Queries) *V1API {
	return &V1API{
		Server:     server,
		Controller: controller,
	}
}
