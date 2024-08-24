package routes

import (
	v1 "danieljonguitud.com/aws-events-go/api/v1"
	"net/http"
)

func RegisterRoutes(v1Api *v1.V1API) {
	routes := http.NewServeMux()

	routes.Handle("/", RegisterAuthRoutes(v1Api.Controller))
	routes.Handle("/events", RegisterEventsRoutes(v1Api.Controller))

	v1Api.Server.Handle("/api/v1/", http.StripPrefix("/api/v1", routes))
}
