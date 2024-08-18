package controllers

import (
	ctx "context"
	"danieljonguitud.com/aws-events-go/db"
	"encoding/json"
	"net/http"
)

func (c *Controller) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.Queries.ListEvents(ctx.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": events,
	})
}

func (c *Controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var eventParams db.CreateEventParams

	err := json.NewDecoder(r.Body).Decode(&eventParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event, err := c.Queries.CreateEvent(ctx.Background(), eventParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": event,
	})
}
