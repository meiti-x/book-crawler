package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// InitializeDatabase initializes the MongoDB connection and returns the collection.
func InitializeDatabase() (*mongo.Client, *mongo.Collection, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection("books")

	return client, collection, nil
}

// DisconnectDatabase safely disconnects from MongoDB.
func DisconnectDatabase(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}
