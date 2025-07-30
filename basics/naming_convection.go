package basics

import "fmt"

// package names should be in lowercase
type EmployeeGoogle struct {
	FirstName string // pascalCase for struct fields
	LastName  string // pascalCase for struct fields
	Age       int    // pascalCase for struct fields
}

func main(){
	//pascalCase
	//eg. UserInfo, UserProfile, UserAccount	
	//struct , enum, interfaces

	//snake_case
	//eg .user_info, user_profile, user_account
	//variables, functions, methods, packages

	//UPPERCASE
	//constant are emphasized  and immutability are emphasized
	//used in constants

	//mixedCase

	const METRICSYSTEM = 5
	
   var  employeeID = 1001
   fmt.Println("Employee ID:", employeeID)

}
