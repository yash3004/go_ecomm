package Users

import (
	"encoding/json"
	"fmt"
	"go_ecomm/Mongodb"
	"io"
	"log"
	"net/http"

	"go_ecomm/Schema"

	"golang.org/x/crypto/bcrypt"
)

// Define the route handler
func welcomeRoutes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the /welcome route!")
}

func createUser(mh *Mongodb.MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request_body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error occured %v", err)
		}

		defer r.Body.Close()
		var user *Schema.User
		if user.UserId == "" {
			response := map[string]string{"error": "please give user_id"}
			jsonData, _ := json.Marshal(response)
			w.Write(jsonData)
		}
		if err := json.Unmarshal(request_body, &user); err != nil {
			log.Printf("error while parsing to the user  %v", err)
		}

		if err := Create(mh, user); err != nil {
			log.Printf("error occured while creating new user %v \n", err)
		}
		response := map[string]bool{"success": true}
		jsonData, _ := json.Marshal(response)
		w.Write(jsonData)
	}
}

func getUser(mh *Mongodb.MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.PathValue("id")
		user, err := Read(mh, user_id)
		read_user := Schema.PublicUser{
			UserId:     user.UserId,
			Username:   user.Username,
			PublicId:   user.PublicId,
			FirstName:  user.FirstName,
			MiddleName: user.MiddleName,
			LastName:   user.LastName,
			Email:      user.Email,
			PhoneNo:    user.PhoneNo,
		}
		if err != nil {
			log.Printf("error while fetching the user %v", err)
		}
		json_data, err := json.Marshal(&read_user)

		if err != nil {
			log.Printf("error while converting to json %v", err)
		}
		w.Write(json_data)

	}
}

func validateUser(mh *Mongodb.MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth_user, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("error occured while fetching the user id password ")
		}
		var user Schema.ValidateUser

		if err := json.Unmarshal(auth_user, &user); err != nil {
			log.Printf("error while parsing validation")
			return
		}
		log.Printf(user.UserId, user.Password, user.UserName)
		saved_user, mongo_err := Read(mh, user.UserId)
		if mongo_err != nil {
			log.Printf("error while getting the user details %v", mongo_err)
			return
		}

		hashErr := bcrypt.CompareHashAndPassword([]byte(saved_user.Password), []byte(user.Password))
		if hashErr != nil {
			response := map[string]string{"login": "failed"}
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(response)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(jsonData)
			return
		}

		// Password matches
		response := map[string]string{"login": "success"}
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

	}
}

func deleteUser(mh *Mongodb.MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		del_user, fetch_err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if fetch_err != nil {
			log.Printf("cant fetch the data ,%v ", fetch_err)
			return
		}
		var user Schema.ValidateUser
		if err := json.Unmarshal(del_user, &user); err != nil {
			log.Printf("cant parse the data to auth user struct %v", err)
			return
		}

		if err := Delete(mh, user.UserId); err != nil {
			log.Printf("cant delete the user %v with the error %v", user.UserId, err)
			response := map[string]string{"delete": "failed"}
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(response)
			w.WriteHeader(http.StatusFailedDependency)
			w.Write(jsonData)
			return
		}

		response := map[string]string{"delete": "success"}
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

	}
}

func addAddr(mh *Mongodb.MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func updateAddr(mh *Mongodb.MongoHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Register the routes
func RegisterUsersRoutes(mh *Mongodb.MongoHandler) {
	http.HandleFunc("/welcome", welcomeRoutes)
	http.HandleFunc("/create_user", createUser(mh))
	http.HandleFunc("/get_user/{id}", getUser(mh))
	http.HandleFunc("/validate_user", validateUser(mh))
	http.HandleFunc("/delete_user", deleteUser(mh))
	http.HandleFunc("/add_address/{id}", addAddr(mh))
	http.HandleFunc("/update_address/{id}", updateAddr(mh))

}
