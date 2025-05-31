package main

import (
	"fmt"
	"runtime"
	"time"
)

func leakyGoroutine() {
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go leakyGoroutine()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Active goroutines:", runtime.NumGoroutine())
}
