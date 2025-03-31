package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define CLI flags
	name := flag.String("name", "User", "Your name")
	age := flag.Int("age", 18, "Your age")
	verbose := flag.Bool("verbose", false, "Enable verbose output")

	// Parse command-line flags
	flag.Parse()

	// Print output based on flags
	if *verbose {
		fmt.Printf("Verbose Mode Enabled\n")
	}
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
