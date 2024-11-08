package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace the URI with your MongoDB connection string
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Ping the primary to verify connection
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	MongoClient = client
	log.Println("Connected to MongoDB!")

	// Create the database if it doesn't exist
	db := client.Database("notes")
	err = db.CreateCollection(ctx, "notes")
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		log.Printf("Failed to create collection: %v", err)
	}

	// Create the plans collection if it doesn't exist
	db = client.Database("plans")
	err = db.CreateCollection(ctx, "plans")
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		log.Printf("Failed to create plans collection: %v", err)
	}

	return nil
}

func GetNotesCollection() *mongo.Collection {
	return MongoClient.Database("notes").Collection("notes")
}

func GetPlansCollection() *mongo.Collection {
	return MongoClient.Database("plans").Collection("plans")
}
