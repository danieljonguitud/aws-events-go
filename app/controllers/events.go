package controllers

import (
	ctx "context"
	"log"
	"net/http"
	"time"

	"danieljonguitud.com/aws-events-go/db"
)

const eventsHTML = "app/views/events/index.html"

func (c *Controller) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.Queries.ListEvents(ctx.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RenderView(w, eventsHTML, events)
}

func (c *Controller) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var userId = int64(1)
	var eventParams db.CreateEventParams

	log.Print("Request received")
	log.Print(r.Header)
	log.Print(r.Header.Get("HX-Request"))

	now := time.Now().UTC()

	eventParams.Name = r.PostFormValue("name")
	eventParams.Description = r.PostFormValue("description")
	eventParams.Location = r.PostFormValue("location")
	eventParams.Datetime = now
	eventParams.UserID = userId

	event, err := c.Queries.CreateEvent(ctx.Background(), eventParams)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RenderFragment(w, eventsHTML, "event-list-element", event)
}
