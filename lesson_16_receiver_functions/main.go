package main

import "fmt"

func main() {
	// Pass a customer name into newBill, which returns a bill object, stored as myBill - a bill instance
	myBill := newBill("Alex")

	fmt.Println(myBill.format()) 
}

// Your Bill: 
// pie: ...$5.99 
// cake: ...$3.99 
// total: ...$9.98
//? we can improve the look and align this... 
// before - fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
//* after - fs += fmt.Sprintf("%v ...$%-25v \n", k+":", v)