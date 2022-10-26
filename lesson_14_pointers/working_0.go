package main

import "fmt"


func updateName(x string) {
	//*  Updating the name inside the function, updates the local copy, not the original value
	x = "wedge"
	//? To combat this we can use pointers. 
}

func main() {
	//? 1. name is passed into updateName function to update string value 
	//* Go creates a copy of name.
	name := "tifa"
	//* copy is passed into function 
	updateName(name)

	fmt.Println(name)
}