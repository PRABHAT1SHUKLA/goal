package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 20
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Clear screen
	fmt.Print("\033[2J")
	// Hide cursor
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h") // Show cursor when done

	columns := make([]int, width)

	for {
		for i := 0; i < width; i++ {
			if columns[i] < height && rand.Float32()
