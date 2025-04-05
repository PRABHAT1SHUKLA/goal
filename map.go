package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	// Add
	colors["blue"] = "#0000ff"

	// Delete
	delete(colors, "green")

	// Iterate
	for key, value := range colors {
		fmt.Println(key, "->", value)
	}
}
