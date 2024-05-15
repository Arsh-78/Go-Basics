package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s Shape) {

	fmt.Println("Area of square is :", s.getArea())
}

func main() {
	s := Square{sideLength: 4}
	t := triangle{base: 5, height: 4}

	printArea(s)
	printArea(t)

}
