package app

import (
	"danieljonguitud.com/aws-events-go/db"
	"net/http"
)

type App struct {
	Server *http.ServeMux
}

func New(server *http.ServeMux, db *db.Queries) *App {
	return &App{
		Server: server,
	}
}
