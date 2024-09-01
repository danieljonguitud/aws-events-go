package controllers

import (
	ctx "context"
	"encoding/json"
	"net/http"

	"danieljonguitud.com/aws-events-go/db"
)

func (c *Controller) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.Queries.ListEvents(ctx.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data": events,
	}

	ResponseHandler(w, http.StatusOK, response)
}

func (c *Controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	userId := int64(1)
	var eventParams db.CreateEventParams

	err := json.NewDecoder(r.Body).Decode(&eventParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventParams.UserID = userId

	event, err := c.Queries.CreateEvent(ctx.Background(), eventParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data": event,
	}

	ResponseHandler(w, http.StatusCreated, response)
}
