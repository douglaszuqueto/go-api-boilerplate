package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger wraps an HTTP handler and logs the request as necessary.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("(%s) \"%s %s %s\" %s", r.RemoteAddr, r.Method, r.RequestURI, r.Proto, time.Since(start))
	})
}
