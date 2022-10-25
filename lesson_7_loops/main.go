package main

import (
	"fmt"
)

func main() {
	//? like a while loop
	// initialise var = 0
	x := 0
	// while x is less than 5
	for x < 5 {
		// print value of x
		fmt.Println("value of x is:", x)
		// each iteration add 1
		x++ // infinite loop without this
	}

	//? 'traditional' for loop - does same thing 
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	// while i is less than length of names slice, cycle though names in list and print each name 
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// cycle though names slice
	// each time get the index and the value 
	for index, val := range names {
		// format string
		fmt.Printf("the value at pos %v is %v \n", index, val)
		// this will not alter value in original slice
		val = "new string"
	}
	
	//? what if we just want to use the value, not index
	for _, val := range names {
		fmt.Print(val, ",")
		// this will not alter value in original slice
		val = "new string"
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)

}