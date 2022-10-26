package main

import "fmt"


func updateName(x string) {
	//* The way we give a pointer to a variable is by placing a  in front of it
	x = "wedge"
	
}

func main() {

	name := "tifa"
	 
	updateName(name)

	//? The & gets us a pointer to the memory location of the name variable
	fmt.Println("memory address of name is:", &name)
	// memory address of name is: 0x14000110210
	//* we can store pointers in a variable - m is store the memory location in its own memory block
	m := &name
	// the address for name, stored in m 
	fmt.Println("memory address:", m)
	//? using an asterisk before the pointer will get the value pointed to - 'tifa'. 
	fmt.Println("value at memory address:", *m) // tifa


	// fmt.Println(name)
}


//* we can store pointers in a variable - m is storing the pointer/memory location  of the name variable in its own memory block 0x002
/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/