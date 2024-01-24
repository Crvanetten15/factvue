package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response struct to hold the JSON response
type Response struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Update this to match your Vue app's URL
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check if the method is "OPTIONS" for preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Define the handler with CORS middleware
	http.Handle("/", corsMiddleware(http.HandlerFunc(jsonHandler)))

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// jsonHandler sends a JSON response
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a response object
	response := Response{
		Message: "Change in data",
		Path:    r.URL.Path,
	}

	// Encode the response as JSON
	json.NewEncoder(w).Encode(response)
}
