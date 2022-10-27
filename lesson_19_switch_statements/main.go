package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// why error? is this returned from ReadString method? 
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	// getInput returns the user input which we store in vars
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		fmt.Println(tip)
	case "s":
		fmt.Println("you chose to save the bill")
	// default fired if a, t or s isn't the opt
	default:
		fmt.Println("That was not a valid option...")
		// recall function
		promptOptions(b)
	}
}

func main() {

	mybill := createBill()
	promptOptions(mybill)

}