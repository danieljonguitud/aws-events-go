package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/app/controllers"
)

func RegisterHomeRoutes(controller *controllers.Controller) http.Handler {
	homeRoutes := http.NewServeMux()

	homeRoutes.HandleFunc("GET /", controller.Index)

	return homeRoutes
}
