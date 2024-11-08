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

func getInteger() (int, error) {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	fmt.Print("시작값 입력: ")
	n1, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("종료값 입력: ")
	n2, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
