package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//todo - create function to hold this logic
//? take in 2 things prompt, string and r - pointer to reader 
//* prompt is the question we want to ask
// return 2 things: 1. string, 2. error
//! don't want to create a new reader each time the function is called
//? so take in a reader as an argument = r pointer to bufio.reader
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	// input = name but could be something else, input is more reusable
	input, err := r.ReadString('\n')
	// trim input string and return
	return strings.TrimSpace(input), err
}

// create bill - return bill object with user input
func createBill() bill {
	//? reader - reads user input
	//* from bufio package - newreader method 
	//? what is the source? standard input, we get this from os package - os.Stdin
	// "i want a reader to read from standard input ie terminal"
	reader := bufio.NewReader(os.Stdin) //* we pass reader into get input


	// fmt.Print("Create a new bill name: ")
	//? use readstr method from the source (os.Stdin), until a new line \n
	// store as name and trim space
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	//todo - create function above to hold this logic
	// take name but not err, arg1 = "Create a new bill name: ", arg2 = reader
	name, _ := getInput("Create a new bill name: ", reader)

	// pass in user input to generate a newBill
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	// local reader
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a -add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {
	// call createBill
	mybill := createBill()
	promptOptions(mybill)

}