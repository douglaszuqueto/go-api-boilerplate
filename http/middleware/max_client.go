package middleware

import "net/http"

var sema = make(chan struct{}, 5)

// MaxClients maxClients
func MaxClients(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sema <- struct{}{}
		next.ServeHTTP(w, r)
		<-sema
	})
}
