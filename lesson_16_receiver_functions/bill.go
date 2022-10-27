package main

import "fmt"

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
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip: 0,
	}
	return b
}

//* format the bill
// return type string
//! func format() string { - without *, the function can be called globally
//? we need to associate format to the bill struct/ bill objects
//* (b bill) - associate with bill, and use a COPY of the instance as b. Can only be called from a bill object... myBill.format()
func (b bill) format() string {
	// fs - formatted str
	fs := "Your Bill: \n"
	var total float64 = 0

	// bill items
	// for key & value on the bill instance passed in, loop the items property and add each key and value to the formatted string
	for k, v := range b.items {
		// Sprintf formats according to a format specifier and returns the resulting string. https://pkg.go.dev/fmt#Sprintf
		//? 1st arg - string 
		//? 2nd arg - variables we want to output
		//* fs += fmt.Sprintf("", k, v)
		//todo - we add to the k a colon, we do it here as we want to add spacing to the key variable in the string later on. 
		// fs += fmt.Sprintf("", k+":", v)
		//* example string - pie: ...$5.99 
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		// add item values to total variable
		total += v
	}

	// total 
	//* example string - total: ...$9.98
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}


// Your Bill: 
// pie: ...$5.99 
// cake: ...$3.99 
// total: ...$9.98
//? we can improve the look and align this... 
// before - fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
//* after - fs += fmt.Sprintf("%v ...$%-25v \n", k+":", v)
//todo - The -25 makes the %v character 25 chars in length, adding extra to the right, pushing the prices to the right, like below

// Your Bill: 
// pie:                      ...$5.99 
// cake:                     ...$3.99 
// total:                    ...$9.98

//* if we'd added the : to the string, and not k
//* - fs += fmt.Sprintf("%v ...$%-25v \n", k+":", v)
// total                    :...$9.98