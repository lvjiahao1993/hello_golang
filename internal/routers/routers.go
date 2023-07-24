// routes.go

package routers

import (
	"database/sql"
	"hello_golang/internal/controllers"
	"net/http"
)

// NewRouter creates a new HTTP router and sets up routes
func NewRouter(db *sql.DB) http.Handler {
	// Create a new HTTP router
	router := http.NewServeMux()

	// Set up routes and their corresponding handlers
	router.HandleFunc("/", controllers.HomeHandler)
	router.HandleFunc("/users", controllers.GetUsersHandler)
	router.HandleFunc("/users/{id}", controllers.GetUserByIDHandler)

	// router.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	// router.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")
	// router.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")

	return router
}
