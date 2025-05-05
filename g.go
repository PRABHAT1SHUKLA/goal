package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Write to a file
	data := []byte("Hello, Golang!")
	err := ioutil.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Read from a file
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Content:", string(content))
}
