package main

import "fmt"

func main() {

//* Ints

var ageOne = 20
var ageTwo = 30
ageThree := 40

fmt.Println(ageOne, ageTwo, ageThree) 
// 20 30 40

//* bits and memory 
// 8 = 8 bits - https://pkg.go.dev/builtin#int8
// int8 is the set of all signed 8-bit integers. Range: -128 through 127.
var numOne int8 = 25
// 215 will throw an error as it is out of the scope of int8, int16 would work.
// var numOne int8 = 215 
// var numTwo int8 = -129

// uint must be a positive number 
var numThree uint = 25
// error
// var numThree uint = -25
//? uint8 - as it cannot be negative you can have a larger positive num - upto 255 
var numThreeAgain uint8 = 25
}

