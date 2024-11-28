package Mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoHandler struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoHandler(mongoURI, dbName, collectionName string) (*MongoHandler, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %v", err)
	}
	database := client.Database(dbName)
	collection := database.Collection(collectionName)
	return &MongoHandler{
		Client:     client,
		Database:   database,
		Collection: collection,
	}, nil
}

func (mh *MongoHandler) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mh.Client.Disconnect(ctx)
}
