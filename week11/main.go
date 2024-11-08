package main

import (
	"fmt"
	"log"
	"week11/keyboard"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			j = j + 2
		}
	}
	return true
}

func main() {
	fmt.Print("시작값 입력: ")
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("종료값 입력: ")
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
