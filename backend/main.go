package main

import (
	"fmt"
	"log"
	"net/http"
)

// Start the server and register routes
func startServer(host string, port int) error {
	// Register routes from routes.go
	RegisterRoutes()

	// Start the server
	return http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), nil)
}

func main() {
	// Start the server on localhost:3000
	err := startServer("localhost", 3000)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
