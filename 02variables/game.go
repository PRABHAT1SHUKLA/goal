package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1

	var guess int
	attempts := 0

	fmt.Println("Welcome to Guess the Number!")
	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		attempts++

		if guess < target {
			fmt.Println("Too low! Try again.")
			fmt.Println("Yes, you're up for it .")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("🎉 Correct! You guessed it in %d attempts.\n", attempts)
			break
		}
	}
}
