package main

import (
	"fmt"

	"math/rand"
)

func main() {
	// rand.Seed(time.Now().Unix())
	// answer := rand.Intn(6) + 1 //dice 1~6
	// reader := bufio.NewReader(os.Stdin)
	// var win bool = false

	// for guesses := 0; guesses < 3; guesses++ {
	// 	fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력: ", 3-guesses)
	// 	input, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	input = strings.TrimSpace(input)
	// 	guess, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(guess)

	// 	if answer == guess {
	// 		fmt.Println("정답입니다.")
	// 		win = true
	// 		break
	// 	} else if answer > guess {
	// 		fmt.Println("입력하신 수는 정답보다 작은 수 입니다. Low")
	// 	} else {
	// 		fmt.Println("입력하신 수는 정답보다 큰 수 입니다. High")
	// 	}
	// }
	// if win == true {
	// 	fmt.Println("당신이 이겼습니다.")
	// } else {
	// 	fmt.Printf("당신이 졌습니다. 정답은 %d입니다.", answer)
	// }

	// fmt.Printf("%f 나누기 %f는 %f입니다", 1.0, 3.0, 1.0/3.0) // console 출력

	result := fmt.Sprintf("%0.2f 나누기 %0.2f는 %0.2f입니다\n", 1.0, 3.0, 1.0/3.0) // 서식을 문자열로 리턴
	fmt.Print(result)

	i := 1
	fmt.Printf("%T\n", i)
	for i <= 10 {
		fmt.Printf("%4d\n", rand.Intn(1000)+1)
		// fmt.Printf("%2d\n", i)
		i++
	}
}
