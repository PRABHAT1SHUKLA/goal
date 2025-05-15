package main

import (
	"fmt"
	"sync"
)

var memo = make(map[int]int)
var mu sync.Mutex

func fibonacci(n int) int {
	mu.Lock()
	val, ok := memo[n]
	mu.Unlock()

	if ok {
		return val
	}
	if n <= 1 {
		return n
	}

	result := fibonacci(n-1) + fibonacci(n-2)

	mu.Lock()
	memo[n] = result
	mu.Unlock()
	return result
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
	}
}
