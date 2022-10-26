package main

import "fmt"

// func updateName(x string) {
// 	x = "yoshi"
// }

func updateName(x string) string {
	x = "wedge"
	return x
}

// passing in a map - string key: float64 value 
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	// group A types will not change the initial value
	//* non-pointer wrapper values
	name := "tifa"

	// updateName(name)
	//? when we pass in a var or value into a function, go creates a copy of it.
	//? so the 'name'  being passed into updateName is a copy, not the original.
	//TODO - to get the variable to change we must reassign it to the returned value of the function - name = updateName(name)
	name = updateName(name)

	fmt.Println(name) // now wedge

	//? group B types -> slices, maps, functions ... arrays?
	//* pointer wrapper values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
	//? here we try to add coffee and its value to the map using the update menu function
	//TODO - this will change the initial value, coffee will be added to the menu
	updateMenu(menu)
	fmt.Println(menu)
}


//*  When we create a variable using a type B, Go does 2 things:
//? 1: It stores the data in its own block in memory (block 1)
//? 2: In another separate block (block 2), Go stores a value that contains other information, including a pointer to block 1. The variable name is associated with block 2.
//todo-  the value is split up, the underlying data and a wrapper that contains a pointer to the underlying data
//todo- When we use this variable, Go finds block 2 containing the pointer, follows the memory address the pointer contains and reads from or writes to that data, memory location.
//* When this variable is passed into a function, Go makes a copy of block 2 - a new variable in memory (block 3), containing the same pointer to block 1. 
//* When we make changes to the variable (block 3), inside the function - Go looks at block 3, sees the pointer to block 1, and updates the block 1 value. Changing the original value.