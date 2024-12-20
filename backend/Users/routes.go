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
		if err := json.Unmarshal(request_body, &user); err != nil {
			log.Printf("Error while parsing the data")
			http.Error(w, "INVALID REQUEST PAYLOAD", http.StatusBadRequest)
		}
		log.Println(user)
		if user.UserId == "" {
			log.Printf("not getting the userid")
			http.Error(w, "USER ID NOT FOUND", http.StatusBadRequest)
			return
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
			http.Error(w, "INVALID REQUEST PAYLOAD", http.StatusBadRequest)
			return
		}
		log.Printf(user.UserId, user.Password, user.UserName)
		saved_user, mongo_err := Read(mh, user.UserId)
		if mongo_err != nil {
			log.Printf("error while getting the user details %v", mongo_err)
			http.Error(w, "DB ERROR", http.StatusConflict)
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
		userId := r.PathValue("id")
		log.Printf("user id : %v", userId)
		raw_addr, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("cant get the user request %v", err)
		}
		var userAddr Schema.UserAddress
		json_Err := json.Unmarshal(raw_addr, &userAddr)
		if json_Err != nil {
			log.Printf("cant parse to json ")
			http.Error(w, "INVALID REQUEST", http.StatusBadRequest)
			return
		}
		fmt.Println(userAddr.Address1)
		userAddr.UserId = userId
		if err := addAddress(mh, &userAddr); err != nil {
			log.Printf("error occur while adding to db")
			http.Error(w, "CANT ADD TO DB", http.StatusConflict)
			return
		}
		log.Printf("Address added successfully")
		response := map[string]string{"process": "success"}
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

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
