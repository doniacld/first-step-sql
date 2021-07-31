package main

import (
	"log"
	"net/http"

	"github.com/doniacld/first-step-sql/internal/handlers"
)

// main runs the service
func main() {
	router := handlers.NewService().Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
