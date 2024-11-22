package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myArray [5]float64 = [5]float64{3.5, 4.1, 4.5, 2.9, 3.7}
	mySlice := myArray[0:3]

	// mySlice[1] = 2.2
	myArray[1] = 2.2
	mySlice = append(mySlice, 4.3, 1.1)

	fmt.Println(len(myArray), myArray, reflect.TypeOf(myArray))
	fmt.Println(len(mySlice), mySlice, reflect.TypeOf(mySlice))
}
