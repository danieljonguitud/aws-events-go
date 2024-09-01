package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/app/controllers"
)

func RegisterEventsRoutes(controller *controllers.Controller) http.Handler {
	eventRoutes := http.NewServeMux()

	eventRoutes.HandleFunc("GET /", controller.GetAllEvents)
	eventRoutes.HandleFunc("POST /", controller.CreateEvent)

	return eventRoutes
}
