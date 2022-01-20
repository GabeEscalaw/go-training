package main

import (
	"fmt"
	"sort"
)

func getMedian(array []int) int {
	if len(array) % 2 != 0 {
		return array[len(array)/2]
	} else {
		midLeft := array[len(array)/2]
		midRight := array[len(array)/2 - 1]
		return (midLeft + midRight)/2
	}
}

func main() {
	var num int
	numArray := make([]int, 0)
	
	_, err := fmt.Scan(&num)

	for err == nil {
		numArray = append(numArray, num)
		sort.Ints(numArray)
		fmt.Println(getMedian(numArray))
		_, err = fmt.Scan(&num)
	}

}

