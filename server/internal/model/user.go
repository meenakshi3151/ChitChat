package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a MongoDB document in the "users" collection
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`  // MongoDB ObjectID
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`  
	Password  string             `bson:"password" json:"password"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at"`
	UserName  string			 `bson:"username" json:"username"`
}
