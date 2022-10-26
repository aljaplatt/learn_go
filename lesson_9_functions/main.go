package main

import (
	"fmt"
	"math"
)
// n is the param and it will be a string
func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

//? if you want to return something you must specify what its type will be. the second float64 is the return tag, the type you want to return 
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleNames([]string{"cloud", "barret", "tifa"}, sayGreeting)
	cycleNames([]string{"cloud", "barret", "tifa"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)
}

