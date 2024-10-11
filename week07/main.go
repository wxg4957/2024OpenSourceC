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
	/* var now time.Time = time.Now()
	var year int = now.Year()
	var month int = int(now.Month())
	var day int = now.Day()
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", year, month, day)
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", now.Year(), int(now.Month)), now.Day())
	fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다.\n", now.Hour(), now.Minute(), now.Second()) */

	/* var army string = "우리는 $가와 $민에 충성을 다하는 대한민$ 육군이다."
	armyFixed := strings.NewReplacer("$", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army)) */

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input your score: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)            //역슬래시 부분 제거, strip과 유사
	score, _ := strconv.ParseInt(input, 16, 32) //16진수 정수형 32비트
	if score >= 60 {
		fmt.Println("A")
		fmt.Println(score)
	} else {
		fmt.Println("BCDF")
	}
}
