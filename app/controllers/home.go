package controllers

import (
	"net/http"
	"text/template"
)

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/views/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
