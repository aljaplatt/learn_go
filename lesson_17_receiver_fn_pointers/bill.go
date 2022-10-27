package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

//* function to generate new bills object 
func newBill(name string) bill {
	
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}

//* add item to the bill
func (b *bill) addItem(item string, price float64) {
	b.items[item] = price
}

//* format the bill
// With structs, pointers are automatically dereferenced, so adding one here makes sense as its cheaper to make a copy of a pointer, than it is the object its self.
func (b *bill) format() string {
	
	fs := "Your Bill: \n"
	var total float64 = 0

	for k, v := range b.items {
		
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//* add tip
	// create formatted string, add to fs. 
	//todo - for b.tip to work, need to add pointer, to ensure og value is updated. not a copy
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// total 
	//* example string - total: ...$9.98
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)

	return fs
}

//* update tip with receiver function
//todo - for b.tip to work, need to add pointer, to ensure og value is updated, not a copy.
//? * indicates pointer to bill - (b *bill)
func (b *bill) updateTip(tip float64) {
	//* when we take in a pointer the argument is automatically dereferenced by Go.  
	// (*b).tip = tip - dereferenced 
	b.tip = tip
}

