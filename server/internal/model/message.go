package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Message represents a chat message in MongoDB
type Message struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`  // Unique Message ID
	Message string             `bson:"message" json:"message"`   // Message content
	From    primitive.ObjectID `bson:"from" json:"from"`         // Sender's User ID
	To      primitive.ObjectID `bson:"to" json:"to"`             // Receiver's User ID
}

