package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("Guess the number (1 to 100):")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Correct! 🎉")
			break
		}
	}
}
