package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"log"
)

// DBClient holds the MongoDB client instance for reusability
var DBClient *mongo.Client

// Db initializes the MongoDB connection and returns a client
func Db() {
	fmt.Println("Database connection file!!")
	LoadEnv()

	// Fetch Mongo URI from environment variables
	dbMONGOURI := os.Getenv("MONGO_URI")
	if dbMONGOURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbMONGOURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	// Store client globally for reusability
	DBClient = client

	// Ping the database to ensure the connection is successful
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

// GetDBClient returns the MongoDB client instance for use in other parts of the application
func GetDBClient() *mongo.Client {
	return DBClient
}

// GetCollection returns a specific MongoDB collection
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return DBClient.Database(databaseName).Collection(collectionName)
}
