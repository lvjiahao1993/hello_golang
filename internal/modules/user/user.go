// user.go

package user

import (
	"errors"
	"fmt"
)

// User represents a user data structure
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// NewUser creates a new User instance with the given ID, username, and email
func NewUser(id int, username, email string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
	}
}

// GetUserByID returns a user by ID
func GetUserByID(userID int) (*User, error) {
	// Logic to fetch the user from a database or other data source
	// For this example, we'll just return a dummy user
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	return NewUser(userID, "john_doe", "john@example.com"), nil
}

// SaveUser saves a user to the database or other data source
func (u *User) SaveUser() error {
	// Logic to save the user to the database or other data source
	// For this example, we'll just print the user's information
	fmt.Printf("Saving user: ID=%d, Username=%s, Email=%s\n", u.ID, u.Username, u.Email)
	return nil
}

// DeleteUser deletes a user from the database or other data source
func DeleteUser(userID int) error {
	// Logic to delete the user from the database or other data source
	// For this example, we'll just print a message indicating deletion
	fmt.Printf("Deleting user with ID=%d\n", userID)
	return nil
}
