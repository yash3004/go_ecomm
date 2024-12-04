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

	_, err = mh.Database.Collection("Users").InsertOne(ctx, u)
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
	err := mh.Database.Collection("Users").FindOne(ctx, bson.M{"userid": _id}).Decode(&result)
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
	_, err := mh.Database.Collection("Users").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}
	fmt.Println("Document updated successfully")
	return nil
}

func Delete(mh *Mongodb.MongoHandler, _id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mh.Database.Collection("Users").DeleteOne(ctx, bson.M{"userid": _id})
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}
	fmt.Println("Document deleted successfully")
	return nil
}

func addAddress(mh *Mongodb.MongoHandler, userAddr UserAddress) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := mh.Database.Collection("UserAddress").InsertOne(ctx, userAddr)
	if err != nil {
		return fmt.Errorf("error while inserting the userAddress ")
	}
	return nil
}

func updateAddress(mh *Mongodb.MongoHandler, user_id string, userAddr UserAddress) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"user_id": user_id}
	update := bson.M{"$set": userAddr}
	_, err := mh.Database.Collection("UserAddress").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error occured while updating user_id %v", user_id)
	}
	return nil
}
