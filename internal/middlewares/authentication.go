// authentication.go

package middlewares

import (
	"net/http"
)

// AuthenticationMiddleware is a middleware to check if the request is authenticated.
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is authenticated
		if !isAuthenticated(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authenticated, call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// isAuthenticated checks if the request is authenticated (dummy implementation).
func isAuthenticated(r *http.Request) bool {
	// Add authentication logic here
	// For simplicity, this is just a dummy implementation
	return true
}
