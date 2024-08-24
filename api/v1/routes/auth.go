package routes

import (
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/controllers"
)

func RegisterAuthRoutes(controller *controllers.Controller) http.Handler {
	usersRoutes := http.NewServeMux()

	usersRoutes.HandleFunc("POST /register", controller.RegisterNewUser)
	usersRoutes.HandleFunc("POST /login", controller.LoginUser)

	return usersRoutes
}
