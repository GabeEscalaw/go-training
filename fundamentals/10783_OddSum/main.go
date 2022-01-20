package main

import "fmt"

func main() {
	var T, a, b, sum int
	fmt.Scan(&T)
	sumArray := make([]int, 0, T)
	for i := 0; i < T; i++ {
		fmt.Scan(&a, &b) 
		for j := a; j < b+1; j++ {
			if j % 2 != 0 {
				sum += j
			}
		}
		sumArray = append(sumArray, sum)
		sum = 0
	}
	for i := 0; i < len(sumArray)+1; i++ {
		fmt.Printf("Case %d: %d\n", i+1, sumArray[i])
	}
}