package routes

import (
	"net/http"
	"text/template"

	"danieljonguitud.com/aws-events-go/app"
)

func RegisterRoutes(app *app.App) {
	routes := http.NewServeMux()

	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("hello").Parse(`
            <!DOCTYPE html>
            <html lang="en">
            <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <title>Hello World</title>
            </head>
            <body>
                <h1>Hello, World!</h1>
            </body>
            </html>
        `))
		tmpl.Execute(w, nil)
	})

	app.Server.Handle("/", routes)
}
