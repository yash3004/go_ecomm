package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Define the route handler
func welcomeRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the /welcome route!")
}

func createUser(mh *MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request_body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error occured %v", err)
		}

		defer r.Body.Close()
		var user User
		if err := json.Unmarshal(request_body, &user); err != nil {
			log.Printf("error while parsing to the user  %v", err)
		}

		if err := mh.Create(user); err != nil {
			log.Printf("error occured while creating new user %v \n", err)
		}
		response := map[string]bool{"success": true}
		jsonData, _ := json.Marshal(response)
		w.Write(jsonData)
	}
}

func getUser(mh *MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request_body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	log.Printf("error occured while fetching the request body %v", err)
		// }
		// defer r.Body.Close()
		// var user_id string
		// if err := json.Unmarshal(request_body, &user_id); err != nil {
		// 	log.Printf("error while parsing the user_id , %v", err)

		// }
		user_id := r.PathValue("id")
		user, err := mh.Read(user_id)
		read_user := ReadUser{user.UserID, user.Username, user.Email, user.PhoneNo}
		if err != nil {
			log.Printf("error while fetching the user %v", err)
		}
		json_data, err := json.Marshal(read_user)

		if err != nil {
			log.Printf("error while converting to json %v", err)
		}
		w.Write(json_data)

	}
}

func validateUser(mh *MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth_user, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error occured while fetching the user id password ")
		}
		var user AuthUser

		if err := json.Unmarshal(auth_user, &user); err != nil {
			log.Printf("error while parsing validation")
		}
		saved_user, mongo_err := mh.Read(*user.UserID)
		if mongo_err != nil {
			log.Printf("error while getting the user details %v", err)
		}

		hash_err := bcrypt.CompareHashAndPassword([]byte(saved_user.Password), []byte(user.Password))
		if hash_err != nil {
			response := map[string]string{"login": "failed"}
			json_data, _ := json.Marshal(response)

			w.Write(json_data)
		}
		response := map[string]string{"login": "success"}
		json_data, _ := json.Marshal(response)
		w.Write(json_data)

	}
}

// Register the routes
func RegisterRoutes(mh *MongoHandler) {
	http.HandleFunc("/welcome", welcomeRoutes)
	http.HandleFunc("/create_user", createUser(mh))
	http.HandleFunc("/get_user/{id}", getUser(mh))
	http.HandleFunc("/validate_user", validateUser(mh))
}
