package controllers

import (
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":   "Home Page",
		"Message": "Events App!",
	}

	RenderView(w, "app/views/index.html", data)
}
