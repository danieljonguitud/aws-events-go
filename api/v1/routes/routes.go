package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/controllers"
)

func RegisterRoutes(controller *controllers.Controller) {
	v1 := http.NewServeMux()

	v1.Handle("/events/", RegisterEventsRoutes("/events", controller))

	controller.Server.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))
}
