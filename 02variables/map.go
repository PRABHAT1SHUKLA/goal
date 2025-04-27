package main

import "fmt"

func main() {
	fruits := map[string]int{
		"apple":  5,
		"banana": 3,
		"mango":  8,
	}

	for fruit, count := range fruits {
		fmt.Printf("We have %d %s(s)\n", count, fruit)
	}
}
