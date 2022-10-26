// this connects greetings to main.go
package main

import "fmt"

var points = []int{20,90,100,45,70}

// usable in main.go
func sayHello(n string) {
	fmt.Println("hello", n)
}

func showScore(){
	fmt.Println("you scored this many points:", score)
}