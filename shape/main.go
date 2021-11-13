package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	s := square{5}
	t := triangle{5, 7}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return t.base * t.height * (1 / 2.0)
}

func printArea(s shape) {
	fmt.Println("The area of the shape is ", s.getArea())
}
