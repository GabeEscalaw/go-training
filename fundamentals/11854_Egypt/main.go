package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	for err == nil {
		if a == 0 && b == 0 && c == 0 {
			return
		} else if ((a*a + b*b) == c*c ) || ((a*a + c*c) == b*b ) || ((b*b + c*c) == a*a ) {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
		
		_, err = fmt.Scan(&a, &b, &c)
	}
}