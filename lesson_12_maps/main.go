package main

import "fmt"

func main() {

	//? keys must be all one data type, values should all be same value
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	// "pie" - key
	fmt.Println(menu["pie"])

	// looping maps
	// k - key, v - value
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[647583927] = "yoshi"
	fmt.Println(phonebook)
}