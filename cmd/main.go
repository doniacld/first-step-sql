package main

import (
	"log"
	"net/http"
	"os"

	"github.com/doniacld/first-step-sql/internal/db/postgres"
	"github.com/doniacld/first-step-sql/internal/handlers"
)

// main runs the service
func main() {
	log.Printf("%s %s %s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := postgres.New(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	if err != nil {
		log.Fatalf("unable to establish connection with Postgres database: %s", err)
	}
	router := handlers.NewService(db).Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
