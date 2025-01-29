package repository

import (
	"ChitChat/internal/config"
	"ChitChat/internal/model"
	"context"
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

func CheckIfExists(user *model.User) string {
	collection := config.GetCollection("ChitChat", "users")
	err := collection.FindOne(context.Background(), user)
	if err != nil {
		return "User Logged in successfully"
	}
	return "User does not exist! Please sign up"
}