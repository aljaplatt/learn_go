package main

//? go run main.go greetings.go - have to run both files for association to work 

import "fmt"

// score has to be declared outside main function to be accessible in other files
var score = 99.5

// cannot use shorthand outside of functions
// scoreTwo := 50

func main() {
	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}
}

