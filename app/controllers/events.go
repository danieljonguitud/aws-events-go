package controllers

import (
	ctx "context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"danieljonguitud.com/aws-events-go/db"
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

	htmlStr := fmt.Sprintf("<ul><li>Name: %s</li><li>Description: %s</li><li>Location: %s</li><li>Creator: %s</li></ul>", event.Name, event.Description, event.Location, event.UserID)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}
