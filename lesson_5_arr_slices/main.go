

package main

import "fmt"

func main() {
	//* array has to have a fixed length that cannot change 
	//? [3] - length, int - datatype = 
	// var ages [3]int = [3]int{20, 25, 30}
	//? type inferred
	//todo - ARRAY - add num/length inside square bracket 
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	//* you can change items... change mario to luigi.
	names[1] = "luigi"

	fmt.Println(ages, len(ages)) 
	// [20 25 30] 3
	fmt.Println(names, len(names))
	// [yoshi luigi peach bowser] 4

	//? slices (use arrays under the hood)
	// more like arrays in other languages - we can add, or take away items 
	//todo - SLICES - do not add num inside square brackets 
	var scores = []int{100, 50, 60}
	// change 60 -> 25
	scores[2] = 25
	// add/append item to slice
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//? slice ranges - get an section of items from an existing arr or slice and store them in a new slice
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	// append to rangeOne
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}