package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	for err == nil {
		if b > a {
			fmt.Println(b - a)
		} else {
			fmt.Println(a - b)
		}
		
		_, err = fmt.Scan(&a, &b)
	}
	
}