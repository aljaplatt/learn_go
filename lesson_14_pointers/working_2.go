package main

import "fmt"

//? we can allow this function to take a pointer to a string by adding an asterisk before the string type declaration 
//* the asterisk in front of a type allows the use of a pointer to the value in that location.
// func updateName(x string) {
func updateName(x *string) {
	// x = "wedge"
	//? with the asterisk in place, we're now updating the block holding the value - dereference operator
	*x = "wedge"
	
}

func main() {

	name := "tifa"
	fmt.Println(name) //! tifa
	
	// pointer - save the memory location to the variable name as the variable m
	m := &name

	//todo - we can pass the pointer into a function 
	updateName(m) // passing in the pointer updates the original value
	fmt.Println(name) //* now wedge 
}


//* we can store pointers in a variable - m is storing the pointer/memory location  of the name variable in its own memory block 0x002
/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/