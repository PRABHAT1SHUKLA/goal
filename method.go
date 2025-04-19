package main

import "fmt"

type Developer struct {
	Name string
	Role string
}

func (d Developer) Introduce() {
	fmt.Printf("Hi, I’m %s and I work as a %s\n", d.Name, d.Role)
}

func main() {
	dev := Developer{"Prabhat", "Golang Developer"}
	dev.Introduce()
}
