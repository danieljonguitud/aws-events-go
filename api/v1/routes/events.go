package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/controllers"
)

func RegisterEventsRoutes(basePath string, controller *controllers.Controller) http.Handler {
	eventRoutes := http.NewServeMux()

	eventRoutes.HandleFunc("GET /", controller.GetAllEvents)

	return http.StripPrefix(basePath, eventRoutes)
}
