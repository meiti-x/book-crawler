package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitializeDatabase initializes the MongoDB connection and returns the collection.
func InitializeDatabase(uri, dbName, collectionName string) (*mongo.Client, *mongo.Collection, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return client, collection, ctx, nil
}

// DisconnectDatabase safely disconnects from MongoDB
func DisconnectDatabase(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}
