package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(writer http.ResponseWriter, statusCode int, message string) {
	if statusCode >= 500 {
		log.Println("Responding with 5XX error:", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(writer, statusCode, errorResponse{
		Error: message,
	})
}

func respondWithJSON(writer http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(data)
}
