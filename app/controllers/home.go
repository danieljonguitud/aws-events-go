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
		"Title":   "Welcome Page",
		"Message": "Hello, World!",
	}

	RenderView(w, "app/views/index.html", data)
}
