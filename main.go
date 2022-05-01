package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func sumAreas(Shapes []Shape) float64 { // Shape is an interface
	total := 0.0
	for _, s := range Shapes {
		total += s.Area()
	}
	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s := Square{20}
	c := Circle{10}
	Shapes := []Shape{s, c}
	fmt.Println(sumAreas(Shapes))
}
