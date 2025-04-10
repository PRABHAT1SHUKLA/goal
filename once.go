package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	init := func() {
		fmt.Println("Initializing only once!")
	}

	for i := 0; i < 5; i++ {
		go func() {
			once.Do(init)
		}()
	}

	time.Sleep(1 * time.Second)
}
