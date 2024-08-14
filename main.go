package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"danieljonguitud.com/aws-events-go/api/v1/controllers"
	v1Routes "danieljonguitud.com/aws-events-go/api/v1/routes"
	"danieljonguitud.com/aws-events-go/db"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
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

	mux := http.NewServeMux()

	controller := controllers.New(queries, mux)

	v1Routes.RegisterRoutes(controller)

	if err := http.ListenAndServe(":8080", controller.Server); err != nil {
		fmt.Println(err.Error())
	}
}
