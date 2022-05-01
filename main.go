package main

import (
	"fmt"
)

type Square struct {
	x    int
	y    int
	side float64
}

func NewSquare(x int, y int, side float64) (*Square, error) {
	if side <= 0 {
		return nil, fmt.Errorf("side must be greater than 0")
	}
	return &Square{x, y, side}, nil
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) move(x int, y int) {
	s.x += x
	s.y += y
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
	s.move(2, 3)
	fmt.Printf("%+v\n", s)
}
