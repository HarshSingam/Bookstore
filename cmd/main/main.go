package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterBookStoreRoutes(r)

	// Start the HTTP server
	port := ":8080" // Use a commonly used port for HTTP or choose another available port
	log.Printf("Server is starting on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
