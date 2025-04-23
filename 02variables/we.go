package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct for a response
type Message struct {
	Text string `json:"text"`
}

// Handler function for the root route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Text: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", helloHandler) // Handle route "/"
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)  // Start the server
}
