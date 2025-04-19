package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	v  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func main() {
	c := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final count:", c.Value())
}
