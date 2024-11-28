package main

import (
	"fmt"
	"log"
	"net/http"

	"go_ecomm/Mongodb"
	"go_ecomm/Users"
)

func startServer(host string, port int) error {
	cfg, config_err := get_config()
	if config_err != nil {
		log.Printf("error while loading the configurations %v ", config_err)
	}
	mh, mongoHandler_err := Mongodb.NewMongoHandler(cfg.MongoConfig, "testdb", "Users")
	if mongoHandler_err != nil {
		log.Printf("error while loading mongodb %v", mongoHandler_err)
	}
	Users.RegisterRoutes(mh)

	return http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), nil)
}

func main() {
	err := startServer("localhost", 3000)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
