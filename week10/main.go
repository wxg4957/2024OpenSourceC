package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	// counts := 0
	var isPrime bool = true // 가독성, 메모리 이득
	j := 2
	for j < n {
		if n%j == 0 {
			// counts++
			isPrime = false // 더하기 연산 제거
		}
		j++
	}
	// if counts == 0 {
	if isPrime { // == 비교 연산 제거
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is not prime number.", n)
	}
}
