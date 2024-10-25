package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 //dice 1~6
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Print("숫자 입력: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다.")
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다. Low")
		} else {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다. High")
		}
	}
}
