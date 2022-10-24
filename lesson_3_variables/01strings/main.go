package main

import "fmt"

func main() {

	//* strings - must be a string, must use double quotes
	//! if we do not use the variable it will throw an error
	// you can explicitly declare that you want the variable to be a string
	var nameOne string = "mario"
	// or not - Go will infer type for us
	var nameTwo = "luigi"
	// here we declare an empty variable that we want to use for a string in the future
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)
	
	//? we can reassign the variable
	nameOne = "peach"
	nameThree = "bowser"
	
	fmt.Println(nameOne, nameTwo, nameThree)

	//? we can also use the shorthand and drop the var
	// however, you cannot use this outside of a function
	nameFour := "yoshi"
	fmt.Println(nameFour)
}
