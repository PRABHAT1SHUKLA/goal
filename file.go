package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
