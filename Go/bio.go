package main

import "fmt"

var name, town string
var age int

func main() {
	name := "moses"
	town := "Nairobi"
	age := 25

	fmt.Printf("hello my name is %s and come from %s and i am %d old", name, town, age)
}
