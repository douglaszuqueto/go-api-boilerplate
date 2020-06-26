package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORS cors
func CORS(routes *mux.Router) http.Handler {
	allowedHeaders := []string{"*"}
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
	allowedOrigins := []string{"*"}

	return handlers.CORS(
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowedMethods(allowedMethods),
		handlers.AllowedOrigins(allowedOrigins))(routes)
}
