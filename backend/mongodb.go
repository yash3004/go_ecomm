package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)



type MongoService interface {
	Create(u User) error
	Read(_id string) (*User, error)
	Update(_id string, updatedUser User) error
	Delete(_id string) error
}

type MongoHandler struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoHandler(mongoURI, dbName, collectionName string) (*MongoHandler, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %v", err)
	}
	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &MongoHandler{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (mh *MongoHandler) Create(u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mh.collection.InsertOne(ctx, u)
	if err != nil {
		return fmt.Errorf("error inserting document: %v", err)
	}
	fmt.Println("Document inserted successfully")
	return nil
}

func (mh *MongoHandler) Read(_id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result User
	err := mh.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error finding document: %v", err)
	}

	return &result, nil
}

func (mh *MongoHandler) Update(_id string, updatedUser User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updatedUser}
	_, err := mh.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}
	fmt.Println("Document updated successfully")
	return nil
}

func (mh *MongoHandler) Delete(_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mh.collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}
	fmt.Println("Document deleted successfully")
	return nil
}