package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextID := intSeq(1001)
	fmt.Println(nextID())
	fmt.Println(nextID())
	anotherNextID := intSeq(100001)
	fmt.Println(anotherNextID())

	x := func(x, y int) int {
		return x += y
	}(4, 5)

	fmt.Println(x)
}