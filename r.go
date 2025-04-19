package main

import (
	"fmt"
	"time"
)

func greet(name string, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Hello " + name
}

func main() {
	ch := make(chan string)
	go greet("Prabhat", ch)
	fmt.Println(<-ch)
}
