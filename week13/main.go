package main

import (
	"fmt"
	"reflect"
)

func main() {
	var emptySlice []bool
	// emptySlice = make([]bool, 5)
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice))

	if len(emptySlice) == 0 {
		emptySlice = append(emptySlice, true)
	}
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice))

	var myArray [5]float64 = [5]float64{3.5, 4.1, 4.5, 2.9, 3.7}
	mySlice := myArray[0:3]

	// mySlice[1] = 2.2
	myArray[1] = 2.2
	mySlice = append(mySlice, 4.3, 1.1)

	fmt.Println(len(myArray), myArray, reflect.TypeOf(myArray))
	fmt.Println(len(mySlice), mySlice, reflect.TypeOf(mySlice))
}
