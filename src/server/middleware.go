package server

import (
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simplified example: In real scenarios, you will be performing security checks, rate limiting etc.
		log.Println("Requested URI:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simplified example: In real scenarios, you might be checking against a database or a JWT token.
		if r.Header.Get("X-Auth-Token") != "my-secret-token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
