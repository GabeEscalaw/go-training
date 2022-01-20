package main

import (
	"fmt"
	"sort"
)

func main() {
	var T, N int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N) 
		runners := make([]int, 0, T)
		var temp int
		for j := 0; j < N; j++ {
			fmt.Scan(&temp)
			runners = append(runners, temp) 
			sort.Ints(runners)	
		}
		fmt.Printf("Case %d: %d\n", i+1, runners[len(runners)-1])
	}
}

