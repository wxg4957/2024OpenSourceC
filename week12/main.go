package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {
	var gpas [3]float64
	for i := 0; i < len(gpas); i++ {
		fmt.Print("Input Float Number: ")
		gpas[i], _ = keyboard.GetFloat()
	}
	for _, gpa := range gpas {
		fmt.Println(gpa)
	}
}
