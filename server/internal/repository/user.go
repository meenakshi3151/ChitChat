package repository

import (
	"context"
	"ChitChat/internal/config"
	"ChitChat/internal/model"
)

// InsertUser inserts a new user into the MongoDB database
func InsertUser(user *model.User) error {
	// Get the MongoDB collection
	collection := config.GetCollection("ChitChat", "users")

	// Insert the user into the collection (MongoDB will internally convert to BSON)
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
