package controllers

import (
	ctx "context"
	"net/http"
)

func (c *Controller) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.Queries.ListEvents(ctx.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Events": events,
	}

	RenderView(w, "app/views/events/index.html", data)
}
