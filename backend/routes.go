package main

import (
	"fmt"
	"net/http"
)

// Define the route handler
func welcomeRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the /welcome route!")
}

// Register the routes
func RegisterRoutes() {
	http.HandleFunc("/welcome", welcomeRoutes)
}
