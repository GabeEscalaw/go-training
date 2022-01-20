package main

import "fmt"

func main() {
	var T, a, b, c int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&a, &b, &c) 
		if a <= 20 && b <= 20 && c <= 20 {
			fmt.Printf("Case %d: %s\n", i+1, "good")
		} else {
			fmt.Printf("Case %d: %s\n", i+1, "bad")
		}
		
	}
}