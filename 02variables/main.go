package main

import "fmt"

//we can't declare  jwt:= 3000000 here we can use it inside main only

func main() {
	var name string = "Prabhat"
	fmt.Println(name)
	fmt.Printf("tupe %T",name)//prints type of name that is string

	var small uint8 = 255 //only value till 255 are in bound for uint
	fmt.Println(small)

	var ran int
	fmt.Println(ran)//go bydefault puts 0 in integer instead of garbage value

	var sr string
	fmt.Println(sr) //nothing by default in string

	//no var style 
    //this walrus method of declaring variable is only allowed inside any method not outside it 
	numbersofpeople := 20002020020 //most common way to declare variable 
	fmt.Println(numbersofpeople)

	ku := 121223
	fmt.Println(ku)
	
}