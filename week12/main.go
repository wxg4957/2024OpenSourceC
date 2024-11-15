package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myArray [5]float64 = [5]float64{3.5, 4.1, 4.5, 2.9, 3.7}
	mySlice := myArray[0:3]
	fmt.Println(myArray, reflect.TypeOf(myArray))
	fmt.Println(mySlice, reflect.TypeOf(mySlice))
}
