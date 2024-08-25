package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateBookIDIndex is for index books based on book_id
func CreateBookIDIndex(ctx context.Context, collection *mongo.Collection) error {
	// Define the index model with "book_id" as the key and unique constraint
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"book_id": 1},            // Index in ascending order
		Options: options.Index().SetUnique(true), // Ensure the index is unique
	}

	// Create the index in the collection
	indexName, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	fmt.Printf("Created index: %s\n", indexName)
	return nil
}
