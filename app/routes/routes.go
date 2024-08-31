package routes

import (
	"danieljonguitud.com/aws-events-go/app"
	"net/http"
)

func RegisterRoutes(app *app.App) {
	routes := http.NewServeMux()

	routes.Handle("/", RegisterHomeRoutes(app.Controller))
	routes.Handle("/events", RegisterEventsRoutes(app.Controller))

	app.Server.Handle("/", routes)
}
