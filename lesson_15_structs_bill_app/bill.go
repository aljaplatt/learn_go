package main

// new type, name, type = struct = structure
// blueprint for any bill object 
// struct is the go equivalent of a class/ data type blueprint
type bill struct {
	// property type
	name string
	// order items - map type, keys strings, values floats 
	items map[string]float64
	tip float64
}

//* function to generate new bills object 
// takes in name as string, return a bill type/object
// make new bills object - call function take in name and assign the bill property - return a bill type
func newBill(name string) bill {
	// b = bill object
	b := bill{
		// define initial values for the object properties
		// name property is set to the name argument passed into the function
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}