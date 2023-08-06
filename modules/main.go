package main

import (
	"fmt"
)

// constants - you cannot change the constant's value once declared
/*
Sample -
	Public Constants
		const LoginToken = "somevalue" - here the first letter of the constant should be  block letter to be used as a public variable
	Private
		const loginToken = "somevalue" - here it is not needed
*/
const LoginToken = "skjiohofiiohnoih"

// function - main
/*
Main entrypoint for us to write the code
*/
func main() {
	/*
	* Basic variable declaration , datatypes and print formats
	 */

	var new_var string = "abhi"
	fmt.Println(new_var)
	fmt.Printf("Variable is of type: %T \n", new_var)

	var isCheckedIn bool = false
	fmt.Println(isCheckedIn)
	fmt.Printf("Variable is of type: %T \n", isCheckedIn)

	var income float64 = 235600.45
	fmt.Println(income)
	fmt.Printf("Variable is of type: %T \n", income)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 256.90877967857656747
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//aliases and default value
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Variable is of type: %T \n", defaultInt)

	var defaultStr string
	fmt.Println(defaultStr)
	fmt.Printf("Variable is of type: %T \n", defaultStr)

	// alternate ways to declare stuff

	var randomNumber = 234.567
	fmt.Println(randomNumber)
	fmt.Printf("Variable is of type: %T \n", randomNumber)

	randomString := "This is a random string !!"
	fmt.Println(randomString)
	fmt.Printf("Variable is of type: %T \n", randomString)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
