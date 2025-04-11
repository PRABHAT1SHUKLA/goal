package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("foo", "bar")
	m.Store("baz", 123)

	m.Range(func(key, value any) bool {
		fmt.Println(key, "=>", value)
		return true
	})

	val, ok := m.Load("foo")
	if ok {
		fmt.Println("Value of foo:", val)
	}
}
