package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	// fmt.Println(scores[1], scores[0]) // ,scores[3])
	// fmt.Printf("%#v\n", scores)
	// fmt.Println(scores)
	for i := 0; i < len(scores); i++ {
		fmt.Printf("%d ", scores[i])
	}

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1947200001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1947200001, 0), // comma
	// }
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1947200001, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])
	// fmt.Printf("%#v\n", dates)
	// fmt.Println(dates)
	fmt.Println()
	// for i, date := range dates {
	// 	fmt.Println(i, date)
	// }
	for _, date := range dates {
		fmt.Println(date)
	}

}
