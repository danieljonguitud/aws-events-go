package routes

import (
	"net/http"

	v1 "danieljonguitud.com/aws-events-go/api/v1"
	"danieljonguitud.com/aws-events-go/shared/middlewares"
)

func RegisterRoutes(v1Api *v1.V1API) {
	routes := http.NewServeMux()

	routes.Handle("/", RegisterAuthRoutes(v1Api.Controller))
	routes.Handle("/events", middlewares.AuthHandler(RegisterEventsRoutes(v1Api.Controller)))

	v1Api.Server.Handle("/api/v1/", http.StripPrefix("/api/v1", routes))
}
