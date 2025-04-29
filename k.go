package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{"John Doe", "john@example.com"}

	// Convert struct to JSON
	jsonData, _ := json.Marshal(user)
	fmt.Println(string(jsonData))

	// Convert JSON back to struct
	var newUser User
	json.Unmarshal(jsonData, &newUser)
	fmt.Println(newUser)
}
