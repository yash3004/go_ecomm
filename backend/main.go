package main

import (
	"fmt"
	"log"
	"net/http"
)

// Start the server and register routes
func startServer(host string, port int) error {
	//load the cfg 
	cfg,config_err := get_config()
	if config_err != nil{
		log.Printf("error while loading the configurations %v " , config_err)
	}
	mh , mongoHandler_err := NewMongoHandler(cfg.MongoConfig , "testdb" , "Users")
	if mongoHandler_err != nil {
		log.Printf("error while loading mongodb %v" , mongoHandler_err)
	}
	RegisterRoutes(mh)

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
