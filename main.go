package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	v1 "danieljonguitud.com/aws-events-go/api/v1"
	v1Controllers "danieljonguitud.com/aws-events-go/api/v1/controllers"
	v1Routes "danieljonguitud.com/aws-events-go/api/v1/routes"
	"danieljonguitud.com/aws-events-go/app"
	"danieljonguitud.com/aws-events-go/app/controllers"
	appRoutes "danieljonguitud.com/aws-events-go/app/routes"
	"danieljonguitud.com/aws-events-go/db"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed sql/schema.sql
var ddl string

func run() (*db.Queries, error) {
	ctx := context.Background()

	conn, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := conn.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := db.New(conn)

	return queries, nil
}

func main() {
	queries, err := run()

	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServeMux()

	apiController := v1Controllers.New(queries)
	v1Api := v1.New(server, apiController, queries)
	v1Routes.RegisterRoutes(v1Api)

	appController := controllers.New(queries)
	app := app.New(server, appController, queries)
	appRoutes.RegisterRoutes(app)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		fmt.Println(err.Error())
	}
}
