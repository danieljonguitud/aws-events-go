package main

import (
	"context"
	v1Routes "danieljonguitud.com/aws-events-go/api/v1/routes"
	"database/sql"
	_ "embed"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

//go:embed schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	v1Routes.RegisterRoutes(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
