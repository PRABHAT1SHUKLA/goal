package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current Time:", now)

	// Add 2 hours
	future := now.Add(2 * time.Hour)
	fmt.Println("After 2 Hours:", future)

	// Format time
	fmt.Println("Formatted:", now.Format("2006-01-02 15:04:05"))
}
