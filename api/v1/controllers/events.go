package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Controller) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	events, err := c.Queries.ListEvents(context.Background())

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
		"events":     events,
	})
}
