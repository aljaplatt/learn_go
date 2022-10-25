package main

import (
	// imports must be on new lines
	"fmt"
	"sort"
	"strings"
)

// these do not replace the original, they return a mew string
func main() {
	greeting := "hello there friends!"

	// see if a string contains 'hello' - returns true or false
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.Contains(greeting, "bye"))

	// replace a term with a new one - replace "hello" -> hi. Returns a new string, old one still exists
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	// turn string to uppercase
	fmt.Println(strings.ToUpper(greeting))

	// returns the position of 'll' = starts at position 2
	fmt.Println(strings.Index(greeting, "ll"))

	// split string into array, empty space " ", means split on empty space
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	//* sort
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	//? sorts ages - will alter original slice
	sort.Ints(ages)
	fmt.Println(ages)
	
	//? int search for 30, returns position - returns 2, position in sorted slice 
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	//? string search  
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	// sorts alphabetically
	sort.Strings(names)
	fmt.Println(names)

	// returns position 0
	fmt.Println(sort.SearchStrings(names, "bowser"))

}