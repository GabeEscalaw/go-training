package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct {
	base, height float64
}

type square struct {
	side float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func (s square) area() float64 {
	return s.side * s.side
}

func main() {
	list := []shape{square{5}, triangle{5, 6}}
	for _, item := range list {
		fmt.Println(item.area())
	}	
}