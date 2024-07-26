package main

import (
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add headers here

		next.ServeHTTP(w, r)
	})
}

// func logRequest(next http.Handler) http.Handler {
