package main

import (
	"fmt"
	"math"
)



type Square struct {
	Side float32
}

func (s *Square) CalculateArea() float32 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float32
}

func (c *Circle) CalculateArea() float32 {
	return math.Pi * c.Radius * c.Radius
}

type AreaCalculator interface {
	CalculateArea() float32
}

func main() {
	square := &Square{Side: 4}
	circle := &Circle{Radius: 4}

	for _, shape := range []AreaCalculator{square, circle} {
		fmt.Println(shape.CalculateArea())
	}
}