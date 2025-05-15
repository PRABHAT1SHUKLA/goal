package main

import (
	"fmt"
	"time"
)

func light(color string, ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("Light is:", color)
			time.Sleep(1 * time.Second)
			ch <- true
			return
		}
	}
}

func main() {
	colors := []string{"Red", "Green", "Yellow"}
	ch := make(chan bool)

	go light(colors[0], ch)
	go light(colors[1], ch)
	go light(colors[2], ch)

	ch <- true // start the first light

	time.Sleep(4 * time.Second) // allow time for all lights
}
