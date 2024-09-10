package main

import "net/http"

// This is a specific function signature as recommended by the Go standard library
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
