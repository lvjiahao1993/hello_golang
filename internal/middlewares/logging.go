// logging.go

package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware to log HTTP requests and their durations.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request and its duration
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}
