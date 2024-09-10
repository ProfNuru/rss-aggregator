package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Error codes in 400s are client side errors
	if code > 499 {
		log.Printf("Responding with %v error: %v", code, msg)
	}

	// json.Marshal requires the struct to indicate corresponding keys in json format
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: msg})
}

// Return JSON formatted repsonse
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal payload into json string and data as bytes
	// We can write bytes directly into http response
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marsha JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
