package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/services"
	"danieljonguitud.com/aws-events-go/auth"
	"danieljonguitud.com/aws-events-go/db"
)

func (c *Controller) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var userParams db.CreateUserParams

	err := json.NewDecoder(r.Body).Decode(&userParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := services.HashPassword(userParams.Password)
	userParams.Password = hashPassword

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := c.Queries.CreateUser(context.Background(), userParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // This should be change to not tell that email should be unique
		return
	}

	token, err := auth.GenerateJWT(user.ID, user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // This should be change to not tell that email should be unique
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	ResponseHandler(w, http.StatusCreated, response)
}

func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userParams db.CreateUserParams

	err := json.NewDecoder(r.Body).Decode(&userParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	token, err := auth.GenerateJWT(user.ID, user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // This should be change to not tell that email should be unique
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	ResponseHandler(w, http.StatusCreated, response)
}
