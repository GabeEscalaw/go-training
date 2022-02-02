package main

import (
	"fmt"
)

// func say(s string) {
// 	for i := 0; i < 3; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say ("world")
// 	go say ("hello")
// }

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total
}

func main() {
	s := []int{8, 4, -3, 10, 2, 1}
	c := make(chan int)
	go sum(s[:3], c)
	go sum(s[3:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

