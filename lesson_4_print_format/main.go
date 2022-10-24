package main

import "fmt"

func main() {
	age := 100
	// age := "100"
	name:= "alex"

	// Print - unlike Println, Print doesnt add on a new line
	fmt.Print("hello, ")
	// fmt.Print("world!")
	fmt.Print("world! \n")
	fmt.Print("new line \n")
	// hello, world! 
	// new line 

	// print variable to terminal - alex is 100 years old
	fmt.Println(name, "is", age, "years old")

	//* Printf - formatted string - order matters - no auto new line
	fmt.Printf("%v is %v years old \n", name, age)
	// with quotes
	fmt.Printf("%q is %q years old \n", name, age)
	//? "alex" is 'd' years old - d didnt work as it was a number, if you change num to str it will work

	// age is of type int 
	fmt.Printf("age is of type %T \n", age)

	// you scored 2.434300 points 
	fmt.Printf("you scored %f points \n", 2.4343)
	// you scored 2.43 points 
	fmt.Printf("you scored %0.2f points \n", 2.4343)

	// Sprintf - save formatted string 
	var str = fmt.Sprintf("%v is %v years old \n", name, age)
	fmt.Println("the saved string is:", str)

	// see more format specifiers here:
	// https://golang.org/pkg/fmt/
}