// package main

//* pointers are a data type themselves and when we create them they're stored in their own memory location.
//? they point to another memory location  
//todo - we can manually create pointers for non pointer values - strings, ints, etc. 
//? WHY? -  

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(m) // using pointer as arg, can pass &name directly instead of "m" as well
	fmt.Println(name)

}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|
*/