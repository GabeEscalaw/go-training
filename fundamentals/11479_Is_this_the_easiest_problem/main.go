package main

import "fmt"

func main() {
	var T, a, b, c int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&a, &b, &c) 
		if ((a + b) > c) && ((a + c) > b) && ((b + c) > a) {
			if (a == b) && (a == c) && (b == c) {
				fmt.Printf("Case %v: Equilateral\n", i+1)
			} else if (a == b) || (a == c) || (b == c) {
				fmt.Printf("Case %v: Isosceles\n", i+1)
			} else {
				fmt.Printf("Case %v: Scalene\n", i+1)
			}
		} else {
			fmt.Printf("Case %v: Invalid\n", i+1)
		}
	}
}