package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	// var isPrime bool = true
	if n <= 1 {
		// isPrime = false
		return false
	} else if n == 2 {
		// isPrime = true
		return true
	} else if n%2 == 0 {
		//isPrime = false
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				// isPrime = false
				// break
				return false
			}
			j = j + 2
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("정수 입력: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if isPrime(n) {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is not prime number.", n)
	}
}
