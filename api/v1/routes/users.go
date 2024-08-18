package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/controllers"
)

func RegisterUsersRoutes(basePath string, controller *controllers.Controller) http.Handler {
	usersRoutes := http.NewServeMux()

	usersRoutes.HandleFunc("POST /", controller.CreateUser)

	return http.StripPrefix(basePath, usersRoutes)
}
