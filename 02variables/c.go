package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Message from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}
