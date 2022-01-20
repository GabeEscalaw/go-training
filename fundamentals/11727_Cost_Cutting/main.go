package main

import (
	"fmt"
	"sort"
)

func main() {
	var T, a, b, c int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&a, &b, &c) 
		salaries := []int{a, b, c}
		sort.Ints(salaries)
		fmt.Printf("Case %d: %d\n", i+1, salaries[1])
	}
}