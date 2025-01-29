package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ChitChat/internal/model"
	"ChitChat/internal/repository"
)

// SignUpUser handles user signup requests
func SignUpUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User Signup handler")

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming JSON request body into the User struct
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Just for testing, print the received user data
	fmt.Printf("Received User Signup Request: %+v\n", user)

	// Insert the user into the database (MongoDB)
	err = repository.InsertUser(&user)
	if err != nil {
		http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
		return
	}

	// Send a response to the client
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User signed up successfully"})
}
