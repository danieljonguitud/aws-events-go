package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/app"
)

func RegisterRoutes(app *app.App) {
	routes := http.NewServeMux()

	routes.Handle("/", RegisterHomeRoutes(app.Controller))
	routes.Handle("/events", RegisterEventsRoutes(app.Controller))

	app.Server.Handle("/", routes)
}
