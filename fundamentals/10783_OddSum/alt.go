package main

import "fmt"

func main() {
	var T, a, b int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&a, &b) 
		sum := 0
		for j := a; j < b+1; j++ {
			if j % 2 != 0 {
				sum += j
			}
		}
		fmt.Printf("Case %v: %v\n", i+1, sum)
	}
}