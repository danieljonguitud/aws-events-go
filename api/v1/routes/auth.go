package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/controllers"
)

func RegisterAuthRoutes(basePath string, controller *controllers.Controller) http.Handler {
	usersRoutes := http.NewServeMux()

	usersRoutes.HandleFunc("POST /register", controller.RegisterNewUser)

	return http.StripPrefix(basePath, usersRoutes)
}
