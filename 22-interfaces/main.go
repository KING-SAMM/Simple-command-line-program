package main

import (
	"fmt"
	"math"
)

// Shape interface
type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	lenght float64
}
type circle struct {
	radius float64
}

// Square methods
func (s square) area() float64 {
	return s.lenght * s.lenght
}
func (s square) perimeter() float64 {
	return s.lenght * 4
}

// Circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Print shape information
func printShapeInfo(s shape) {
	fmt.Printf("Area of %T = %0.2f \n", s, s.area())
	fmt.Printf("Perimeter of %T = %0.2f \n", s, s.perimeter())
}

func main() {
	shapes := []shape{
		square{lenght: 2.5},
		circle{radius: 7.0},
		circle{radius: 1.4},
		square{lenght: 12.2},
	}

	// Loop through shapes and print shape info
	for _, v := range shapes{
		printShapeInfo(v)
		fmt.Println("----")
	}
}