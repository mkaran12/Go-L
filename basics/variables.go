package basics

import "fmt"

var middleName = "Michael"

func main() {
	age := 22
	name := "John Doe"
	isStudent := true
	println("Name:", name)
	println("Age:", age)
	println("Is Student:", isStudent)
	println("Age in months:", age*12)
	println("Is Student:", isStudent)

	println("my name is", name, "and my age is", age, "and I am a student:", isStudent)

	fmt.Println(middleName)

	//default values
	//numeric values
	//Boolen type
	// strings values
	// pointers slices, functions, maps, structs: nil

	// ------scope------
	
}
func firstname() {
	firstname := "John"
	fmt.Println(firstname)
}
 