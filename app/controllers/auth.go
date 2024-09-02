package controllers

import (
	"context"
	"log"
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/services"
	"danieljonguitud.com/aws-events-go/db"
)

func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("login request")
	var userParams db.CreateUserParams

	userParams.Email = r.PostFormValue("email")
	userParams.Password = r.PostFormValue("password")

	user, err := c.Queries.GetUser(context.Background(), userParams.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isValidPassword := services.CheckPasswordHash(user.Password, userParams.Password)

	if !isValidPassword {
		http.Error(w, "Failed to validate credentials", http.StatusUnauthorized)
		return
	}

	//token, err := auth.GenerateJWT(user.ID, user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // This should be change to not tell that email should be unique
		return
	}

	http.Redirect(w, r, "/events", http.StatusSeeOther)
}
