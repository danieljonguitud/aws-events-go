package app

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/app/controllers"
	"danieljonguitud.com/aws-events-go/db"
)

type App struct {
	Server     *http.ServeMux
	Controller *controllers.Controller
}

func New(server *http.ServeMux, controller *controllers.Controller, db *db.Queries) *App {
	return &App{
		Server:     server,
		Controller: controller,
	}
}
