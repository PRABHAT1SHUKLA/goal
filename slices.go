package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)

	for i, v := range numbers {
		fmt.Printf("Index %d: Value %d\n", i, v)
	}
}
