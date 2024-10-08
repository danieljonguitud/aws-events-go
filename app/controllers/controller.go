package controllers

import (
	"net/http"
	"text/template"

	"danieljonguitud.com/aws-events-go/db"
)

type Controller struct {
	Queries *db.Queries
}

func New(queries *db.Queries) *Controller {
	return &Controller{
		Queries: queries,
	}
}

func RenderView(w http.ResponseWriter, file string, data any) {
	tmpl := template.Must(template.ParseFiles(file))

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}

func RenderFragment(w http.ResponseWriter, file string, fragment string, data any) {
	tmpl := template.Must(template.ParseFiles(file))

	if err := tmpl.ExecuteTemplate(w, fragment, data); err != nil {
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
