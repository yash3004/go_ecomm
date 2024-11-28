package Users

import (
	"context"
	"fmt"
	"log"
	"time"

	"go_ecomm/Mongodb"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func Create(mh *Mongodb.MongoHandler, u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error occurred while hashing the password: %v", err)
		return err
	}
	u.Password = string(hashedPass)

	_, err = mh.Collection.InsertOne(ctx, u)
	if err != nil {
		return fmt.Errorf("error inserting document: %v", err)
	}
	fmt.Println("Document inserted successfully")
	return nil
}

func Read(mh *Mongodb.MongoHandler, _id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result User
	err := mh.Collection.FindOne(ctx, bson.M{"userid": _id}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error finding document: %v", err)
	}

	return &result, nil
}

func Update(mh *Mongodb.MongoHandler, _id string, updatedUser User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userid": _id}
	update := bson.M{"$set": updatedUser}
	_, err := mh.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}
	fmt.Println("Document updated successfully")
	return nil
}

func Delete(mh *Mongodb.MongoHandler, _id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mh.Collection.DeleteOne(ctx, bson.M{"userid": _id})
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}
	fmt.Println("Document deleted successfully")
	return nil
}
