// controllers.go

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HomeHandler handles requests to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our website!")
}

// GetUsersHandler handles requests to get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch all users from the database or other data source
	// For this example, we'll just return a list of dummy users
	users := []string{"User 1", "User 2", "User 3"}

	// Convert users slice to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByIDHandler handles requests to get a user by ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch a user by ID from the database or other data source
	// For this example, we'll just return a dummy user with the given ID
	userID := r.URL.Query().Get("id")
	user := fmt.Sprintf("User with ID %s", userID)

	// Write the user information to the response
	fmt.Fprintln(w, user)
}

// CreateUserHandler handles requests to create a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to create a new user based on the request data
	// For this example, we'll just print a message indicating user creation
	fmt.Fprintln(w, "Creating a new user...")
}

// UpdateUserHandler handles requests to update a user by ID
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to update a user based on the request data and ID
	// For this example, we'll just print a message indicating user update
	fmt.Fprintln(w, "Updating a user...")
}

// DeleteUserHandler handles requests to delete a user by ID
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Logic to delete a user based on the given ID
	// For this example, we'll just print a message indicating user deletion
	fmt.Fprintln(w, "Deleting a user...")
}
