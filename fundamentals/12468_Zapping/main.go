package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	_, err := fmt.Scan(&a, &b)
	for err == nil {
		if (a == -1 && b == -1) {
			break
		} 
		
		c = (100 - a) - (100 - b) 
		// This is done to contain the values within 100. 
		// Pos = Channel A < Channel B 
		// (e.g. 19 and 42)
		// Neg = Channel B < Channel A 
		// (e.g. 99 and 38)

		if (c >= 0 && c < 50)  {
			fmt.Println(c) 
			// Channel is pos where we're already on the proposed channel or less than 50 clicks if going up 
			// (e.g. 19 42 | 23 up vs 67 down)  
		} else if (100 - math.Abs(c)) > 50 {
			fmt.Println(math.Abs(c)) 
			// Channel is neg where it's less than 50 clicks down 
			// (e.g. 81 33 | 52 up vs 48 down)
		} else {
			fmt.Println(100 - math.Abs(c)) 
			// Channel is pos, but it's less than 50 going down 
			// (e.g. 12 74 | 62 up vs 28 down)
			// or Channel is neg, but it's less than 50 going up 
			// (e.g. 99 38 | 61 down vs 39 up)
		}
		
		_, err = fmt.Scan(&a, &b)
	}
}
