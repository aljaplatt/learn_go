
//? all files that are part of program will specify main
package main

//? importing fmt, a package from the Go standard library
// fmt is for formatting strings and printing messages to the console
import "fmt"

//? this function is the entry point of our application
// go will look for main, if it is called something else - Go will not automatically call the function
// there should only be one main() function in the application
func main() {
	// Println is a method from the fmt package that prints a line to the console
	fmt.Println("Hello, world!")
}

//? Run the file in the terminal with - go run main.go