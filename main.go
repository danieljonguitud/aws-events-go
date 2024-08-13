package main

import (
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
}
