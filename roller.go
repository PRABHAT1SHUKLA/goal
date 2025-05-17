package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1 // random number between 1 and 6
}

func main() {
	fmt.Println("Rolling the dice...")
	fmt.Println("You got:", rollDice())
}
