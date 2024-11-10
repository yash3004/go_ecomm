package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define the route handler
func welcomeRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the /welcome route!")
}


func createUser(mh *MongoHandler) http.HandlerFunc{
	return func (w http.ResponseWriter , r *http.Request){
		request_body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error occured %v", err)
		}
		defer r.Body.Close()
		var user User
		if err := json.Unmarshal(request_body, &user); err != nil {
			log.Printf("error while parsing to the user  %v", err)
		}
		//declaring a mongo handler
		
		
		
		if err := mh.Create(user); err != nil {
			log.Printf("error occured while creating new user %v \n", err)
		}
		w.Write([]byte("hello"))
}

}

// Register the routes
func RegisterRoutes(mh *MongoHandler) {
	http.HandleFunc("/welcome", welcomeRoutes)
	http.HandleFunc("/create_user", createUser(mh))
}
